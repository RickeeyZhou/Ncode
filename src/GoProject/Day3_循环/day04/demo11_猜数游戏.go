package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//1.产生一个随机数
	rand.Seed(time.Now().UnixNano())
	//num1:=rand.Intn(100) // [0,100)
	/*
	[3,15]
	[0,12] + 3
	*13

	[24,89]
	[0,65]+24

	[m,n]
	rand.Intn(n-m+1)+m
	 */

	//fmt.Println(rand.Intn(13) + 3)
	num1 := rand.Intn(100)+1 // [1,100]
	fmt.Println("随机数已经产生：[1,100]")

}
