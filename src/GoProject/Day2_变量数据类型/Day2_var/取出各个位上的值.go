package main

import "fmt"

func main(){
	const a= 1976
	var Qian int
	var Bai int
	var Shi int
	var Ge int
	Qian =a/1000
	Bai =(a%1000)/100
	Shi =(a%1000)%100/10
	Ge=(a%1000)%100%10
	fmt.Println("千位=",Qian)
	fmt.Println("百位=",Bai)
	fmt.Println("十位=",Shi)
	fmt.Println("个位=",Ge)
}
