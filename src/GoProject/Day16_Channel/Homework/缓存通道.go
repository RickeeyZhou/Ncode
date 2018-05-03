package main

import (
	"fmt"
	"strconv"
)

func main(){
	ch1:=make(chan int)
	fmt.Println(len(ch1),cap(ch1))

	ch2:=make(chan int ,5)
	ch2<-1
	fmt.Println(len(ch2),cap(ch2))

	fmt.Println("...................................")
	ch3:=make(chan string,5)
	go sendData3(ch3)
	for v:=range ch3{
		fmt.Println("这是通道输出的第",v,"个数据")
	}
	fmt.Println("main over")

}
func sendData3(fuck chan string ){
for i:=0;i<10;i++{
	fuck<-"data"+strconv.Itoa(i)
	fmt.Println("zi goroutine 像通道写入第, ",i,"个数据")
}
close(fuck)
}
