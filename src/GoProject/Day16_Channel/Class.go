package main

import (
	"fmt"
	"time"
)

func main() {
	//var ch1 chan bool
	ch1:=make(chan bool)
	go func() {
		for i := 100; i > 0; i-- {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
		ch1<-true
	}()
	<-ch1
}
