package parser

import (
	"kongyixueyuan.com/learngo/0504-parser/engine"
	"kongyixueyuan.com/learngo/0504-parser/model"
	"regexp"
	"fmt"
	"strconv"
)

var nameRegexp = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
var ageRegexp = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marrRegexp  = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
func UserInfoParser(contents []byte) (result engine.RequestResult) {

	userInfo := model.UserInfo{}
	// 姓名
	data := nameRegexp.FindSubmatch(contents)
	fmt.Printf("data:%s\n", data)
	// data[1]
	if len(data) >= 2 {
		userInfo.Name = string(data[1])
	}

	// 年龄
	data = ageRegexp.FindSubmatch(contents)
	if len(data) >= 2 {
		age, err := strconv.Atoi(string(data[1]))
		if err == nil {
			userInfo.Age = age
		}
	}

	// 婚姻
	data = marrRegexp.FindSubmatch(contents)
	if len(data) >= 2 {
		userInfo.Marr = string(data[1])
	}
	fmt.Printf("个人信息：%s\n",userInfo)

	return
}
