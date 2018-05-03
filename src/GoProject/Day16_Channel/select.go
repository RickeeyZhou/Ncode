package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "cccc"
	}()
	go func() {
		ch2 <- "dddd"
	}()
	for i := 0; i <= 1; i++ {
		select {
		case <-ch1:
			fmt.Println("this is ch1")
		case <-ch2:
			fmt.Println("rhi is ch2")
		}
	}
}
