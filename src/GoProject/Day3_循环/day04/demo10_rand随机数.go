package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	/*
	随机数：
	 */
	 num1 := rand.Intn(100) //[0,100)
	 fmt.Println(num1)

	 num2 := rand.Float64() // [0.0,1.0)
	 fmt.Println(num2)
	 //[0.0,10.0)
	 fmt.Println(rand.Float64()*10)
	 //[5.0,10.0)
	 //[0,5) + 5-->[5,10)
	 fmt.Println(rand.Float64()*5+5)
	 // seed ,种子：64位的整数：int64

	 s1:=rand.NewSource(42)//根据指定的int类型数值作为种子数，创建了一个资源Source对象
	 r1 := rand.New(s1)//根据s1资源对象获取rand对象r1
	 num3:=r1.Intn(100) // [0,100)
	 fmt.Println(num3)

	 s2 := rand.NewSource(42)
	 r2 := rand.New(s2)
	 num4 := r2.Intn(100)
	 fmt.Println(num4)

	t1:=time.Now() // 获取当前的时间
	fmt.Println(t1)
	//时间戳：当前的时间，距离1970年1月1日0点0分0秒的秒值，纳秒值。
	timeStamp1:=t1.Unix()
	timeStamp2:=t1.UnixNano()
	fmt.Println(timeStamp1)
	fmt.Println(timeStamp2)
	//设置默认rand的种子数
	rand.Seed(timeStamp2)
	num5:=rand.Intn(100)
	fmt.Println(num5)

}


