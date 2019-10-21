package goapp

import "github.com/go-xorm/xorm"

type Engine = xorm.Engine

func NewEngine(driverName string, dataSourceName string, showSQL bool, maxIdleConns int, maxOpenConns int) (*Engine, error) {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	engine.SetMaxIdleConns(maxIdleConns)
	engine.SetMaxOpenConns(maxOpenConns)
	engine.ShowExecTime(showSQL)
	engine.ShowSQL(showSQL)
	return engine, err
}
