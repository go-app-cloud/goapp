package goapp

import (
	"encoding/xml"
	"io/ioutil"
)

/**
 * 加载 XML 配置文件
 */
func LoadXMLConfig(config string, c interface{}) error {
	data, err := ioutil.ReadFile(config)
	if err != nil {
		return err
	}
	if err := xml.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}
