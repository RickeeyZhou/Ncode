package main

import (
	"fmt"
	"math"
)

func main()  {
	/*
	题目：某个公司采用公用电话传递数据，数据是四位的整数，
	在传递过程中是加密的，
	加密规则如下：每位数字都加上5,除以10的余数代替该数字，
	再将第一位和第四位交换，第二位和第三位交换。
	6549---->4901
	11 10 9 14
	1	0 9 4
	4	9 0 1
	 */

	 //fmt.Println(math.Pow(2,3)) // 2的3次方
	 //fmt.Println(math.Pow(10,4)) // 10的4次方
	 //fmt.Println(math.Pow10(6))


	num1:=6549
	arr :=[4]int{}
	//arr[0] = num1 % 10000 / 1000
	//arr[1] = num1 % 1000 / 100 // 6549-->549-->5
	//arr[2] = num1 % 100 / 10 // 6549--》49--》4
	//arr[3] = num1 % 10 / 1
	for i:=0;i<len(arr);i++{
		m := int(math.Pow10(4-i))
		n := int(math.Pow10(3-i))
		arr[i] = num1 % m / n
	}
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		//arr[i] += 5 // arr[i] = arr[i]+5
		//arr[i] %= 10
		arr[i] = (arr[i]+5)%10
	}
	for i:=len(arr)-1;i>=0;i--{
		fmt.Print(arr[i])
	}

	 }
