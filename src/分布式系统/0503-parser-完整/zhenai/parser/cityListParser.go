package parser

import (
	"regexp"
	"分布式系统/0503-parser-完整/engine"
)

const regexpString = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>].+>([^<].+)</a>`

// 城市列表解析器
func CityListParser(b []byte) (r engine.RequestResult){
	// ^[a-z]
	// [^[a-z]]
	// [a-z0-9]+
	//
	// class=""
	// >
	//[^<].+ 汉字城市名字
	match := regexp.MustCompile(regexpString)

	// 查找所有的匹配的字符串
	bytes := match.FindAllSubmatch(b, -1)
	
	for _, m := range bytes {
		//fmt.Printf("City:%s  URL:%s\n", m[2], m[1])
		//fmt.Println(r)
		r.Items = append(r.Items,m[2])
		r.R = append(r.R,engine.Request{
			string(m[1]),
			engine.NilParser,
		})
	}
	return
}
