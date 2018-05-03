package engine

import (
	"fmt"
	"kongyixueyuan.com/learngo/0503-parser-01/fetcher"
)

// 通过引擎去调用,参数

// 种子
//Request{"http://www.zhenai.com/zhenghun",CityListParser}

func Run(r Request) {

	all, err := fetcher.FetchData(r.Url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)

	r.ParserFunc(all)
}
