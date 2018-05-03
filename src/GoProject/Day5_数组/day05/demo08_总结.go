package main

import ("fmt"
	"math/rand"
	"time")

func main()  {
	/*
	随机数：math/rand-->
		Intn(n),-->获取[0,n)的随机整数
		Float64() -->获取[0.0,1.0)的随机小数
	随机数，都叫伪随机数。通过算法算来的。
		种子：int64
	 */
	 num1:=rand.Intn(100) //获取一个随机整数：[0,99]
	 fmt.Println(num1)
	 rand.Seed(time.Now().UnixNano()) //int64
	 num2 := rand.Intn(100)
	 fmt.Println(num2)


	 /*
	 时间戳：当前时间，距离1970年1月1日0点0分0秒0毫秒...
	 1970-1-1 0:0:0.000
	 2018-4-2 16:48:33.786
	 秒：
	 毫秒：
	 微秒：
	 纳秒：
	  */





	arr := [10]int{}
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		num :=rand.Intn(100)+1
		arr[i] = num
	}
	fmt.Println(arr)
}
