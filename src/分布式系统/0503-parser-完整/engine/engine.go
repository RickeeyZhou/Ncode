package engine

import (
	"log"
	"fmt"
	"分布式系统/0503-parser-完整/fetcher"
)

// 通过引擎去调用,参数

// 种子
//Request{"http://www.zhenai.com/zhenghun",CityListParser}

func Run(r ...Request) {   // r 是Request类型切片

	for len(r) > 0 {
		firstRequest := r[0] // 第一个网址
		r = r[1:]            // 其余的网址
		fmt.Println("即将请求：",firstRequest.Url)
		all, err := fetcher.FetchData(firstRequest.Url)
						     // 该网页的每一个字
		if err != nil {
			log.Printf("%v\n", err)
		}
		//fmt.Printf("%s\n", all)

		result := firstRequest.ParserFunc(all)
							 // 解析网址每一个字	
		//fmt.Printf("%s", result)
		r = append(r, result.R...)
	}

}
