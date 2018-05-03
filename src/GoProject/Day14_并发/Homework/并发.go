package main

import (

	"fmt"
	"time"

)

func main(){
	go hello()
	fmt.Println("i am main function ")
	time.Sleep(3*time.Second)

}
func hello(){
	fmt.Println("hello,,i am in another goroutine working")
}
