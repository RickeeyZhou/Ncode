package main

// 动态判断编码
// go get golang.org/x/net/html

import (
	"fmt"
	"kongyixueyuan.com/learngo/0503-fetcher/fetcher"
	"regexp"
)

const cityListURL = "http://www.zhenai.com/zhenghun"

func main() {

	all, err := fetcher.FetchData(cityListURL)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)

	printCityInfo(all)

}

func printCityInfo(s []byte) {
	// ^[a-z]
	// [^[a-z]]
	// [a-z0-9]+
	//
	// class=""
	// >
	//[^<].+ 汉字城市名字
	match := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>].+>([^<].+)</a>`)

	// 查找所有的匹配的字符串
	bytes := match.FindAllSubmatch(s, -1)

	fmt.Printf("%s\n", bytes)

	for _, m := range bytes {
		fmt.Printf("City:%s  URL:%s\n", m[2], m[1])
	}

}
