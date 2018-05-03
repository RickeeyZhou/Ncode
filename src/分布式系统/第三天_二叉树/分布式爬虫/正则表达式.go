package main

import (
	"regexp"
	"fmt"
)
const mail="rickeeyzhou@gamil.co.cn"
func main(){

	math:=regexp.MustCompile(`[a-zA-Z0-9]+@([a-zA-z]+)(\.[a-zA-Z]+){1,2}`)
//创建查找gmail.com的任务
	isExist:=math.MatchString(mail) // 启动任务中的验证函数
	println(isExist)

	s:=math.FindString(mail)
	println(s)

	sArr:= math.FindAllString(mail,-1)
	fmt.Printf("%s",sArr)

// 	sArray:=math.FindAllSubmatch([]byte(mail),-1) //根据功能去补充正则表达式
// 	println(sArray)
// }

// //黎跃春
// //
// <a href="http://www.zhenai.com/zhenghun/yaan"
// class="">雅安</a>
						
//<a href="http://www.zhenai.com/zhenghun/yanan"
// 									class="">延安</a>
						
// 								<a href="http://www.zhenai.com/zhenghun/yanbian1"
// 									class="">延边</a>
						
// 								<a href="http://www.zhenai.com/zhenghun/yancheng"
// 									class="">盐城</a>
						
// 								<a href="http://www.zhenai.com/zhenghun/yangjiang"