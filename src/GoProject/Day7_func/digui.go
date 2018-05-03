package main

import (
	"fmt"
)

func main(){

	var n,m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Println(digui(n))
	fmt.Println(Fre(m))
}
func digui (n int)int{
	if n==1{
		return 1
	}
	return digui(n-1)*n
}

func Fre(i int)int{
	if i==3{
		return 2
	}
	return Fre(i-1)+(i-2)
}
