package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main(){
	t1:=time.Now()
	t2:=t1.Add(time.Nanosecond)
	fmt.Println(t1.Add(1))
	fmt.Println(t1,"t1")
	fmt.Println(t2)

	t3:=t1.Add(time.Minute)
	fmt.Println(t3,".....t3")

	t4:=t1.Add(time.Hour*24)
	fmt.Println(t4)

	t5:=t1.AddDate(0,0,3)
	fmt.Println(t5," ....t5.")

	d1:=t1.Sub(t5)
	fmt.Println(d1)
	fmt.Println(t3.Sub(t1))

	fmt.Println(t1.After(t5))

	time.Sleep(time.Nanosecond)

	rand.Seed(time.Now().UnixNano())
	num:=rand.Intn(10)+1
	fmt.Println(num)
	time.Sleep(time.Duration(num)*time.Second)
	fmt.Println("sleep over")
}
