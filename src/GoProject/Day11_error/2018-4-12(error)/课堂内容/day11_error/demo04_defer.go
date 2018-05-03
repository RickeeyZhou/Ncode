package main

import "fmt"

func main(){// 程序的入口
/*
defer，延迟，等到其他的程序执行结束才执行
	defer函数用于最后执行。

defer函数有多个，按照后进先出的规则：
	-->栈的结构：
	先延迟的后执行
 */
	printString("hello")
	defer printString("王二狗")
	printString("么么哒")
	defer printString("world")
}
func printString(s string)  {
	fmt.Println(s)
}
