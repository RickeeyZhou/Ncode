package main

import (
	"fmt"
	"time"
)

func main(){
	ch:=make(chan int ,1000)
	defer close(ch)
	 go func (){
		  for i:=0;i<10;i++{
			ch<-i
		}
		//close(ch)
	}()
	go func() {
		for {
			a,ok:=<-ch
			if !ok{
				fmt.Println("close")
				return
			}
			fmt.Println("a:",a)
		}
	}()

	fmt.Println("ok")
	time.Sleep(3*time.Second)
}
// 如果将close 放在主函数,就会在传递值的过程中,打断,提前关闭通道.