package main

import "fmt"

func main() {

	ch1 := make(chan string)
	done := make(chan bool)
	go test11(ch1, done)
	data:=<-ch1
fmt.Println("zi ",data)

	<-done
	fmt.Println("over")

}
func test11(fuck chan string, done chan bool) {

	fuck <- "i am xiaoming"
	data := <-fuck
	fmt.Println(fuck)
	fmt.Println("main goroutine ", data)
	done<-true

}
