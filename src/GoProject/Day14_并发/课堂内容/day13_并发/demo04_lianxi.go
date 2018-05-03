package main

import (
	"fmt"
	"time"
)

func main()  {//主goroutine，执行main()
	/*
	子线程1，每睡眠250毫秒打印数字，一共打印5个。1-5
	子线程2，每睡眠400毫秒打印字母，一共打印5个。A-E
	主线程中，启动两个子线程，然后进入睡眠3000毫秒。
	观察程序运行结果。
	*/
	go printNum() // 子groutine，执行printNum()
	go printLetter() // 子goroutine，执行printLetter()
	time.Sleep(3000*time.Millisecond)
	fmt.Println("main，，over。。")
}
func printNum()  {
	for i:=1;i<=5;i++{
		time.Sleep(150*time.Millisecond)
		fmt.Print(i,"\t")
	}
}
func printLetter()  {
	for i:='A';i<='E';i++{
		time.Sleep(400*time.Millisecond)
		fmt.Printf("%c\t",i)
	}
}
