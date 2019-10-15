package goapp

import "github.com/kataras/iris"

type Context = iris.Context
type Map = iris.Map
type Application = iris.Application

func Default() *Application {
	return iris.Default()
}

func Addr(addr string) iris.Runner {
	return iris.Addr(addr)
}
