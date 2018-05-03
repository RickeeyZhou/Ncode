package main

import "fmt"

func main(){
	a:=199
	p1:=&a
	*p1++
	fmt.Println(a,*p1)
}
