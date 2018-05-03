package main

import (
	"time"
	"fmt"
)

/*
把老师代码复习一遍，回顾知识点，串起来
把包理解的更深一点，把作业做了
预习多进程
 */
func mian(){
	t1:=time.Now()
	fmt.Println(t1)

	///get the day
	year,month,day:=t1.Date()
	fmt.Println("年 月 日:",year,month,day)

	hour,minute,second:=t1.Clock()
	fmt.Println("hour , min , s ", hour,minute,second)
	fmt.Println(t1.ISOWeek())
	fmt.Println(t1.YearDay())
}