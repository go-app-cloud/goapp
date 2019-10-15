package goapp

import "github.com/go-xorm/xorm"

type Engine = xorm.Engine

func NewEngine(driverName string, dataSourceName string) (*Engine, error) {
	return xorm.NewEngine(driverName, dataSourceName)
}
