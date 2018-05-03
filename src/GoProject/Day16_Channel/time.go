package main

import (
	"time"
	"fmt"
)

func main(){
slectTest()
}
func slectTest(){
	ch1:=time.After(3*time.Second)
	fmt.Printf("%T\n",ch1)
	fmt.Println(time.Now())
	time2:=<-ch1
	fmt.Println(time2)

	timer:=time.NewTimer(5*time.Second)
	fmt.Printf("%T\n",timer)
	ch2:=timer.C
	fmt.Println(<-ch2)
}