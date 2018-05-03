package engine

import (
	"分布式系统/深度爬虫/fetcher"
	"fmt"
)

//通过引擎调用, 参数

// 种子

func Run(r Request){
		all ,err:=fetcher.FetchData(r.Url)
		if err!=nil{
			panic (err)
		}
		fmt.Printf("%s\n",all)
		r.ParseFunc(all)
}

