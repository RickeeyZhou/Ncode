package parser

import (
	"regexp"
	"fmt"
)
const regexpString= `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>].+>([^<].+)</a>`

func CityListParser(b []byte){
	match := regexp.MustCompile(regexpString)

	// 查找所有的匹配的字符串
	bytes := match.FindAllSubmatch(b, -1)

	fmt.Printf("%s\n", bytes)

	for _, m := range bytes {
		fmt.Printf("City:%s  URL:%s\n", m[2], m[1])
	}

}
