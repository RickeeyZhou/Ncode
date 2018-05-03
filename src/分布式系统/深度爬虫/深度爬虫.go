package main

import (
	"分布式系统/深度爬虫/engine"
	"分布式系统/深度爬虫/zhenai/parser"
)

//const url ="http://www.zhenai.com/zhenhun"
//func main(){
//	all,err:=fetcher.FetchData(url)
//	if  err!=nil{
//		panic(err)
//
//	}
//	fmt.Printf("all",all)
//
//}


func main() {

engine.Run(engine.Request{
	"http://www.zhenai.com/zhenghun",
	parser.CityListParser,
})

}

