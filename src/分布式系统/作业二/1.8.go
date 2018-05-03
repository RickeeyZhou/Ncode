package main

import (
	"fmt"
	"time"
)

func main(){
	abc:=make (chan int,100)
	defer close(abc)
	for i:=0;i<10;i++{
		abc<-i
	}
	go func (){
	for {
		a := <-abc
		fmt.Println("a:", a)
	}
	}()


	fmt.Println("close ")
	time.Sleep(time.Second*100)
}