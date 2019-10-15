package goapp

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestNewEngine(t *testing.T) {
	engine, err := NewEngine("mysql", "root:123456@/mysql?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	_ = engine.Ping()
}
