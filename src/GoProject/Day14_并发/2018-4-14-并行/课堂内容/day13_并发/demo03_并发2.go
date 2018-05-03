package main

import (
	"fmt"
	"time"
)

func main()  {
	go test1()
	go test2()

	time.Sleep(5*time.Second)
}
//打印数字
func test1()  {
	for i:=1;i<=10000;i++{
		fmt.Println("test1函数中i:。。。", i)
	}
}
//打印数字
func test2()  {
	for j:=1;j<=10000;j++{
		fmt.Println("\ttest2函数中j：",j)
	}
}
/*
1.主线程中打印100个数字，子线程中打印100个数字

2.
	子线程1，每睡眠150毫秒打印数字，一共打印5个。1-5
	子线程2，每睡眠400毫秒打印字母，一共打印5个。A-E
	主线程中，启动两个子线程，然后进入睡眠3000毫秒。
	观察程序运行结果。

 */
