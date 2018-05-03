package main

// 动态判断编码
// go get golang.org/x/net/html

import (
	"kongyixueyuan.com/learngo/0503-parser-01/engine"
	"kongyixueyuan.com/learngo/0503-parser-01/zhenai/parser"
)

func main() {
	//无限深入的去爬虫
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.CityListParser,
	})

}


