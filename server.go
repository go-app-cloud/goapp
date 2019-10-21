package goapp

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
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
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

const (
	Success        = 0
	CheckTypeError = -1
	DBError        = -2
)
