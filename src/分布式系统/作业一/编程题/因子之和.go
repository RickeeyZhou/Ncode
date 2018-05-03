package main

import "fmt"


func yinzi (n int) []int{
	yin:=make([]int,0)
	for i:=1;i<n;i++{
		if n%i==0{
			yin =append(yin,i)
		}
	}
	return yin
}
func main(){
	for i:=1;i<=1000;i++{
		yin:=yinzi(i)
		var sum int
		for _,v:=range yin{
			sum =sum+v
		}
		if sum==i{
			fmt.Printf("%d是完美数\n",i)
		}
	}
}