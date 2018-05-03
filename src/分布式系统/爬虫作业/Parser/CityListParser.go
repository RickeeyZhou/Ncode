package Parser

import (

	"分布式系统/爬虫作业/Drive"
	"regexp"

)
const regexpString= `<a href="(.*?)"\s*class="">(.*?)</a>`
// 解析所有数据中的 城市信息
func CityListParser(all []byte)(r Drive.RequestResult){
	match:=regexp.MustCompile(regexpString)

	bytes:=match.FindAllSubmatch(all,-1)

	for _,m:=range bytes{
		r.Items=append(r.Items,m[2])
		r.R=append(r.R,Drive.Request{
			string(m[1]),
			Drive.NilParser,
		})
	}

return
}
