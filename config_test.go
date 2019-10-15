package goapp

import (
	"encoding/xml"
	"log"
	"testing"
)

type config struct {
	XMLName xml.Name `xml:"config"`
	App     app      `xml:"app"`
}
type app struct {
	App  xml.Name `xml:"app"`
	Port int      `xml:"port,attr"`
}

func TestLoadXMLConfig(t *testing.T) {
	c := config{}
	if err := LoadXMLConfig(`./config.conf.xml`, &c); err != nil {
		log.Fatal(err)
	}
	log.Println(c)
}
