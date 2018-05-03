package main

import "fmt"

func main(){
	i1:="A"
	fmt.Println(i1)
	s1:="hellloooworld"
	for i:=0;i<len(s1);i++{
		fmt.Printf("%c,%T",s1[i],s1[i])
	}
	fmt.Println()
	for _,v:=range s1{
		defer fmt.Printf("%c",v)
	}
}
