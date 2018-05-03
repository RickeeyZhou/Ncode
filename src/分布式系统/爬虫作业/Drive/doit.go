package Drive

import (
	"fmt"
	"分布式系统/深度爬虫/fetcher"
	"log"
)

func Doit(r ...Request) {
	//爬取数据
	// 第一步 装换成[]byte
	// 第二步 解析出城市 和 url

	for len(r) > 0 {
		firstRequest:=r[0]
		r=r[1:]
		fmt.Println("即将处理请求:",firstRequest.Url)
		all, err := fetcher.FetchData(firstRequest.Url)

		if err!=nil{
			log.Printf("%v\n",err)
			result:=firstRequest.Parserfunc(all)

			r= append(r,result.R...)
		}

	}
}