package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scanln(&a)
	switch {
	case a > 0:
		fmt.Println("zhengshu")
	case a < 0:
		fmt.Println("fushu")
	default:
		fmt.Println(0)
	}
	fmt.Println("inter  u like month")
	month:=0
	fmt.Scan(&month)
	switch  {
	case  month%2==0&&month!=2 :
			fmt.Println("30")
	case month%2!=0:
		fmt.Println("31")
	default:
		fmt.Println("28")

	}
	switch month{
	case 1,3,5,7,9,11:
		fmt.Println(31)
	case 2:
		fmt.Println(28)
	case 4,6,8,10,12:
		fmt.Println(30)
	default:
		fmt.Println("u are not in earth")
		}
}
