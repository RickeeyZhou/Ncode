package main

import "fmt"

func main(){
	var a,b,c=1,2,3
	var arr2[3]*int
	arr2[0]=&a
	arr2[1]=&b
	arr2[2]=&c
	fmt.Printf("%p",&arr2[0])
}