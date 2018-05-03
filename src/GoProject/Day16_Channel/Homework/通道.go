package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	done := make(chan bool)

	go printnum(done)
	go printletter(done)
	<-done
	<-done

	fmt.Println("main over...")
}
func printletter(done chan bool) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Println("zi goroutine1,printi", i)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	}
	done <- true
}
func printnum(fuck chan bool) {
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 100; j++ {
		fmt.Println("zi goroutine2,print j", j)
	}
	fuck <- true
}


