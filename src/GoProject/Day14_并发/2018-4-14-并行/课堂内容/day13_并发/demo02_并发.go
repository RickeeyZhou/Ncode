package main

import (
	"fmt"
	"time"
)

func main()  {
	/*
	一个goroutine，打印100个数字，另一个goroutine，打印100个字母。
	思路上：创建并启动一个goroutine，指定要执行的函数即可。。

	并发执行goroutine，主的goroutine不能先死。
	 */
	go hello()
	fmt.Println("我是主函数。。。")
	time.Sleep(3*time.Second)

}
func hello()  {
	fmt.Println("hello，，我是在另一条goroutine中执行。。")
}
