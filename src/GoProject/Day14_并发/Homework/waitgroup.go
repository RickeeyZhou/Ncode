package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup
func main(){
	wg.Add(2)
	go tes1()
	go tes2()
	fmt.Println("进入阻塞状态，等待wg的子goroutine结束")
	wg.Wait()
	fmt.Print("main,defuse block")
}
func tes1(){
	for i:=1;i<=1000;i++{
		fmt.Println("tes1 ",i)
	}
	wg.Done()
}
func tes2(){
	defer wg.Done()
	for j:=1;j<=1000;j++{
		fmt.Println("\ttes2 ",j)
	}

}