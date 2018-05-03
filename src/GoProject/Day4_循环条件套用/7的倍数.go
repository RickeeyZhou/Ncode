package main

import "fmt"

func main() {
	//输出7的倍数[1，100]
	for a:=7;a<200;a+=7{
		fmt.Println(a)
	}
	//统计1到5月的总天数
	var sum int
	const (
		a1=30
		a3
		a5
	)
	const a2=28
	const a4=31
	sum=a1+a2+a3+a4+a5
	fmt.Print("day",sum)
	// 用计算机统计 奇数月的总天数
	//循环12次，遇到奇数 加31
	day:=0
	for i:=12;i>0;i--{
		if i%2!=0 {
			day += 31
		}
	}
	fmt.Print("\n",day)
}
