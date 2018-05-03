package main

import (
	"分布式系统/0503-parser-完整/engine"
	"分布式系统/0503-parser-完整/zhenai/parser"
)

// 动态判断编码
// go get golang.org/x/net/html

func main() {
	//无限深入的去爬虫
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.CityListParser,
	})
}


