package main

import (
	"regexp"
	"fmt"
)

// 声明一个需要匹配的字符串
const text  = `
	my email is liyuechun2009@163.com.cn.com
	939442932@qq.com
	liyuechun2009@126.com
	liyuechun@cldy.org
`

// liyuechun2009@163.com

//xxxx@xxx.xxx.xx
// .xxx.xx

func main()  {

	// 传入正则表达式字符串
	// . 代表任意字符
	// .+ 任何多个字符
	// [a-zA-Z0-9] 任意一个小写字母或者大写字母或者一个数字
	// \n 换行
	// \t
	// \.不存在
	//`o{1,3}`
	// fooooood  ooo
	// bc{1,2}  abcabcabab bc bcc
	// bcccbccccc
	// (bc+){1,2}
	// (\.[a-zA-Z]+){1,2} .任意多的字母 这个整体最多能够重复两次，但至少得有一次
	match := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+(\.[a-zA-Z]+){1,2}`)
	// 判断`liyuechun2009@163.com`是否存在于text里面
	isExist := match.MatchString(text)
	fmt.Println(isExist)

	// 返回正则表达式匹配的字符串
	s := match.FindString(text)
	fmt.Println(s)

	// -1 代表所有
	sArr := match.FindAllString(text,-1)
	fmt.Printf("%s",sArr)



}
