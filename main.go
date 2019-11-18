package main

import (
	"fmt"
	"io/ioutil"
)

// 解析文件
func parseFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		return
	}
	fmt.Printf("反序列化成功  conf: %#v\n  port: %#v\n", conf, conf.ServerConf.Port)

}

func parseFile2(filename string)  {
	// 有一些假数据
	var conf Config
	conf.ServerConf.Ip="127.0.0.1"
	conf.ServerConf.Port=8000
	conf.MysqlConf.Port=9000
	err := MarshalFile(filename,conf)
	if err != nil{
		return
	}
}

func main() {
	//parseFile("D:/config.ini")
	parseFile2("D:/my2.ini")
}
