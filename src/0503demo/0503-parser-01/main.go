package main

// 动态判断编码
// go get golang.org/x/net/html

import (
	"fmt"
	"kongyixueyuan.com/learngo/0503-parser-01/fetcher"
	"kongyixueyuan.com/learngo/0503-parser-01/zhenai/parser"
)

const cityListURL = "http://www.zhenai.com/zhenghun"

func main() {

	all, err := fetcher.FetchData(cityListURL)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)

	parser.CityListParser(all)

}


