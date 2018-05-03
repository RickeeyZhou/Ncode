package main

import (
	"fmt"
	"sync"
	"time"
)
var wait sync.WaitGroup
func main(){
	wait.Add(2)

	go pintabc()
	go pintnum()
	wait.Wait()
}
func pintnum() {
	for i:=1;i<=100;i++{
		fmt.Println(i)
		time.Sleep(3*time.Microsecond)
	}
	wait.Done()
}
func pintabc() {
	for i:='A';i<='z';i++{
		fmt.Printf("%q",i)
		time.Sleep(55*time.Millisecond)
	}
	wait.Done()
}