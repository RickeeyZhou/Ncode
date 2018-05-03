package main

import (
	"time"
	"fmt"
)

func main(){
	ch1:=make(chan int)
	ch2:=make(chan int)
	go func (){
		time.Sleep(3*time.Second)
		ch2<-200
	}()
	go func (){
		time.Sleep(3*time.Second)
		ch1<-100

	}()
	select {
		case num1:=<-ch1:
			fmt.Println("ch1中取数据。。",num1)
		case num2:=<-ch2:
			fmt.Println("ch2 data",num2)
	}
}