package main

import (
	"time"
	"fmt"
)

func main(){
	ch1:=time.After(3*time.Second)
	fmt.Printf("%T\n",ch1)
	fmt.Println(time.Now())
	time2:=<-ch1
	fmt.Println(time2)

	time:=time.NewTimer(5*time.Second)
	fmt.Printf("%T\n",time)
	ch2:=time.C
	fmt.Println(<-ch2)



}