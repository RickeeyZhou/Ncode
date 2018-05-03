package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main()  {
	/*
	Duration,时间间隔
	 */
	 t1 := time.Now()
	 t2 := t1.Add(time.Nanosecond)//在当前时间上，累加1纳秒
	 fmt.Println(t1.Add(1))
	 fmt.Println(t1)
	 fmt.Println(t2)

	 t3 := t1.Add(time.Minute)// 当前时间上，累加1分钟
	 fmt.Println(t3)

	 t4:=t1.Add(time.Hour * 24)//当前时间，类加1天
	 fmt.Println(t4)

	t5:=t1.AddDate(0,0,3)//累加：年，月，日
	fmt.Println(t5)

	//两个time之间的间隔
	d1:= t1.Sub(t5) //-72h0m0s
	fmt.Println(d1)
	fmt.Println(t4.Sub(t1)) //24h0m0s
	fmt.Println(t3.Sub(t1)) //1m0s

	fmt.Println(t1.After(t5)) //false
	fmt.Println(t1.Before(t5)) //true


	 //睡眠
	 //time.Sleep(time.Nanosecond)//睡眠1纳秒
	 //time.Sleep(5 * time.Second) //当前正在执行的程序进入睡眠。。
	 //fmt.Println("over....")

	 //睡眠，随机数：[1-10]s
	 // [m,n], rand.Intn(n-m+1)+m
	 rand.Seed(time.Now().UnixNano())
	 num := rand.Intn(10) + 1 //
	 fmt.Println(num)
	 time.Sleep(time.Duration(num) * time.Second)
	 fmt.Println("睡眠结束。。。")


}
