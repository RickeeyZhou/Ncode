package main

import "fmt"

func main() {
	for i := 101; i <= 2000; i++ {
	i:
		for j := 2; j < i; j++ {
			if i%j == 0 {
				break i
			}
			if j==i-1{
				fmt.Print(i,"\t")
			}
		}

	}
}
