package main

import (
	"time"
	"fmt"
)

func main(){
	ch1:=make(chan int)
	go sendData2(ch1)
	for value:=range ch1{
		fmt.Println("read the data",value) //valeu,ok:=<-ch1
	}
fmt.Println("main...over..")
}
func sendData2(fuck chan int){
	for i:=1;i<=10;i++{
		time.Sleep(time.Second)
		fuck<-i
	}
	close(fuck)  ///

}
