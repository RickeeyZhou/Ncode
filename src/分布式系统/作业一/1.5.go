package main

import (
	"runtime"
	"fmt"
)

func main(){
	runtime.GOMAXPROCS(1)
	int_chan:=make(chan int ,1)
	string_chan:=make(chan string ,1)
	int_chan<-1
	string_chan<-"hello"
	fmt.Println("test")
	select {
	case value := <- int_chan :
		fmt.Println(value)
	case value :=<-string_chan:
		panic(value)
	}
}
/*
程序中make(chan int ,1 )的声明,是一个有缓存的通道,像通道写入不会马上堵塞
第一次遇到符合条件的case
执行完就出来了,不会panic
除非 12 13 行交换位置
 */