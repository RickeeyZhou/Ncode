package main

import "fmt"

func main() {
tan()
math()
}
//大家都是一样的
func tan(){
	var sum int=1

	a:=5
	for ;a>=1;a--{
		sum=sum*a
	}
	fmt.Println(sum)
}
func math() {
	for i := 0; i <5; i++ {
		for j:=0;j<5-i;j++  {
			fmt.Print(" ")
		}
		for j:=0;j<2*i+1;j++{
			fmt.Print("*")
		}
		fmt.Println()

	}
}