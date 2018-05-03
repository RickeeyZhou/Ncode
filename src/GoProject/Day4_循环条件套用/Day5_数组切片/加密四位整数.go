package main

import "fmt"

func main() {
	var num int
	fmt.Println("你要加密什么数字？")
	fmt.Scanln(&num)
	var s[4] int
	var a,b,c,d int
	a=num/1000
	b=num%1000/100
	c=num%1000%100/10
	d=num%1000%100%10
	s=[4]int{a,b,c,d}
	for i:=0;i<len(s);i++{
		s[i]+=5
		s[i]=s[i]%10
	}
	s[0],s[2]=s[2],s[0]
	s[1],s[3]=s[3],s[1]
	for i:=0;i<len(s);i++{
		fmt.Print(s[i],"\t")
	}
}
