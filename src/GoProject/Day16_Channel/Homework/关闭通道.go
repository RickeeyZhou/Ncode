package main

import (
	"fmt"
	"time"
)

func main(){
	ch1:=make(chan int)
	go sendData(ch1)

	for {
		num,ok:=<-ch1
		if !ok{
			break
		}
		fmt.Println(num)
		time.Sleep(time.Millisecond)
	}


}
func sendData(fuck chan int ){
	for i:=1;i<=1088;i++{
		fuck<-i
	}
	close(fuck)
}

