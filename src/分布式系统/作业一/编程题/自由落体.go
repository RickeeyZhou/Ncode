package main

import "fmt"

func main(){
	var len int
	high:=100
	for i:=10;i>=1;i--{
		len=len+high
		high=high/2
		fmt.Println(high)
	}
	high=50
	var up int
	for i:=9;i>=1;i--{
		up=up+high
		high=high/2
	}
	len=len+up
	fmt.Println(len)
}
