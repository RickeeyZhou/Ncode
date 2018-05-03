package main

import (
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(1)
	cha1:=make(chan int,1)
	cha1<-1

	fmt.Println("test go")


	<-cha1
}
