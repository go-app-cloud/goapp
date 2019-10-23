package goapp

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	Secret string
}

type App struct {
	AppId     string
	SecretKey string
}
type StandardClaims = jwt.StandardClaims
type Claims struct {
	Data interface{}
	StandardClaims
}

func (p *Token) Build(data Claims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	data.ExpiresAt = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token.Claims = data
	tokenString, err := token.SignedString([]byte(p.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func (p *Token) Parse(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(p.Secret), nil
	})
	if cs, ok := t.Claims.(*Claims); ok && t.Valid {
		return cs, nil
	}
	return nil, err
}
