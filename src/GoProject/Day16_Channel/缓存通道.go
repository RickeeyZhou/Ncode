package main

import (
	"strconv"
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
func main(){

	wg.Add(1)
	cha:=make(chan string ,5)
	go read(cha)
	go write(cha)
	wg.Wait()
}
func write (cha chan string  ){
	for i:=1;i<5;i-- {
		cha <-"data"+strconv.Itoa(i)
		fmt.Println("子goruntine,写出第",i,"个数据")
	}
	close(cha)
}
func read(cha chan string  ){
	for{
		//time.Sleep(time.Second)
		time.Sleep(100*time.Millisecond)
		v,ok:=<-cha
		if !ok{
			fmt.Println("over...",ok)
			break
		}
		fmt.Println("读到数据：",v)
	}
	fmt.Println("main over")
	wg.Done()
}