package main

import "fmt"

func main(){
	a:=[]int{1,2,3,4}
	fmt.Println(a[1],a[2])
	b1:=a
	fmt.Println(b1[1],b1[2])
	fmt.Printf("%p,%p\n",b1,a)
	fmt.Printf("%p,%p",b1[1],a[1])



}
