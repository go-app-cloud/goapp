package goapp

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Context = iris.Context
type Party = iris.Party
type DirOptions = iris.DirOptions
type Map = iris.Map
type Application = iris.Application

func Default() *Application {
	return iris.Default()
}

func Addr(addr string) iris.Runner {
	return iris.Addr(addr)
}

func HTML(directory, extension string) *view.HTMLEngine {
	return view.HTML(directory, extension)
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

const (
	Success        = 0
	CheckTypeError = -1
	DBError        = -2
)

const (
	timeout            = 10
	errorAuthConnect   = `服务登录连接认证失败.`
	successAuthConnect = "服务登录连接认证成功."
)

type AuthRequest struct {
	AppId     string `json:"appid"`
	SecretKey string `json:"secretkey"`
}
type Message struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}
type Socket struct {
	upgrade websocket.Upgrader
	Devices sync.Map
	onAuth  func(request AuthRequest) error
	read    func(conn *websocket.Conn, message Message) error
	close   func(appId string)
	connect func(req AuthRequest, conn *websocket.Conn)
	timer   func(appId string, conn *websocket.Conn)
}

func BuildSocket(auth func(request AuthRequest) error, connect func(req AuthRequest, conn *websocket.Conn), read func(conn *websocket.Conn, message Message) error, close func(appId string), timer func(appId string, conn *websocket.Conn)) *Socket {
	so := Socket{
		upgrade: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		onAuth:  auth,
		close:   close,
		read:    read,
		connect: connect,
		timer:   timer,
	}
	return &so
}

func (p *Socket) SocketHandler(ctx Context) {
	c, err := p.upgrade.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	if err := c.SetReadDeadline(time.Now().Add(time.Second * timeout)); err != nil {
		log.Print("im set read deadline:", err)
		return
	}
	req := AuthRequest{}
	if err := c.ReadJSON(&req); err != nil {
		log.Println(err)
		_ = c.WriteJSON(Response{Code: -1, Msg: errorAuthConnect})
		return
	}
	if err := p.onAuth(req); err != nil {
		_ = c.WriteJSON(Response{Code: -1, Msg: err.Error()})
		return
	}
	defer func() {
		p.Devices.Delete(req.AppId)
	}()
	v, ok := p.Devices.Load(req.AppId)
	if ok {
		cc := v.(*websocket.Conn)
		cc.Close()
	}
	p.Devices.Store(req.AppId, c)
	if err := c.WriteJSON(Response{Code: 100, Msg: successAuthConnect}); err != nil {
		log.Println(err)
		goto ErrorClose
	}
	p.connect(req, c)
	if err := c.SetReadDeadline(time.Time{}); err != nil {
		log.Print("im set read deadline:", err)
		goto ErrorClose
	}
	go func() {
		for {
			if c == nil {
				return
			}
			p.timer(req.AppId, c)
			<-time.After(time.Second * 15)
		}
	}()
	for {
		var msg Message
		if err := c.ReadJSON(&msg); err != nil {
			goto ErrorClose
		}
		if err := p.read(c, msg); err != nil {
			log.Println(err)
			goto ErrorClose
		}
	}
ErrorClose:
	p.close(req.AppId)
}

type SocketClient struct {
	uri    *url.URL
	close  func(socket *SocketClient, uri *url.URL)
	read   func(message json.RawMessage)
	before func(con *websocket.Conn)
	timer  func(con *websocket.Conn) error
	conn   *websocket.Conn
}

func BuildSocketClient(uri *url.URL, read func(message json.RawMessage), before func(con *websocket.Conn), close func(socket *SocketClient, uri *url.URL), timer func(con *websocket.Conn) error) *SocketClient {
	so := SocketClient{
		uri:    uri,
		before: before,
		close:  close,
		read:   read,
		timer:  timer,
	}
	so.ReConnect()
	return &so
}
func (p *SocketClient) ReConnect() {
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		c, _, err := websocket.DefaultDialer.Dial(p.uri.String(), nil)
		if err != nil {
			log.Println("dial:", err)
			goto ErrorClose
		} else {
			p.conn = c
			defer p.conn.Close()
			// pre handler
			p.before(c)
			raw := json.RawMessage{}
			go func() {
				for {
					if c == nil {
						return
					}
					if err = p.timer(c); err != nil {
						return
					}
					<-time.After(time.Second * 15)
				}
			}()
			for {
				if err := p.conn.ReadJSON(&raw); err != nil {
					log.Println("read:", err)
					goto ErrorClose
				}
				p.read(raw)
			}
		}
	ErrorClose:
		p.close(p, p.uri)
	}()
}
func (p *SocketClient) SendJSON(data interface{}) error {
	return p.conn.WriteJSON(data)
}
