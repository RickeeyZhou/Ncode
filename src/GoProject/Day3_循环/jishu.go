package main

import (
	"fmt"
)

func main() {
	for flag:=0,i=1;;i+=2{
		fmt.Print(i)
		flag++
		if flag==20 {
			break
		}

	}
}

