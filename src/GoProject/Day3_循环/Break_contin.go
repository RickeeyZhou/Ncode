package main

import "fmt"

func main() {c := 0 //局部变量
	for i := 0; i <= 100; i++ {

		if i%3 == 0 && i%5 != 0 {
			fmt.Println(i, "\n")
			c++
			if c == 5 {
				break
			}
		}
	}
}
