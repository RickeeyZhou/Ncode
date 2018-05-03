package main

import "fmt"

func main() {
	var month int8
	fmt.Scanln(&month)
	if month%2 != 0 {
		fmt.Println(31)

	} else if month == 2 {
		fmt.Println("runnian")
	}else{
		fmt.Println('s')
	}
}
// 基本数据类型 ： 字符串，数值型，布尔型，