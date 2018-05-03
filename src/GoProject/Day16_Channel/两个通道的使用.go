package main

import (
	"fmt"
	"time"
)

var ch2 = make(chan bool)
var ch1 = make(chan bool)

func main() {

	go PrintAny(1)
	time.Sleep(time.Second)
	go PrintAny('A')

	<-ch1
	<-ch2

}

func PrintAny(byte2 byte) {
	for count:=1;count<=20;count++ {
		if byte2<49 {
			fmt.Println(byte2)
			byte2++
			time.Sleep(time.Second)
			if count==10{
				ch1<-true
			}
		}else{
			fmt.Printf("%c\n",byte2)
			byte2++
			time.Sleep(time.Second)
			if count==10{
				ch2<-true
			}
		}
	}

}
