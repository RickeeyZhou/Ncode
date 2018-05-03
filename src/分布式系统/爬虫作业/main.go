package main

import (
	"分布式系统/爬虫作业/Drive"

	"分布式系统/爬虫作业/Parser"
)

const url ="http://www.zhenai.com/zhenhun"
func main(){
	//一个函数执行url 
	// 执行函数
	Drive.Doit (Drive.Request{
		"http://www.zhenai.com/zhenghun",
		Parser.CityListParser,

	})
	
	
}
