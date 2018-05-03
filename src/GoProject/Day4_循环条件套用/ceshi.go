package main

import "fmt"

func main() {
	var a int

	//var B=300+a
	//var b=3+a*10
	for b := 999; b >= 100; b-- { //循环899次
		B := float64((b-3)/10) + float64((b-3)%10)*0.1 + 300 //通过a 建立B和b的关系   84 的情况384 843
		//B := float64((b-3)/10) + float64((b-3)%10)*0.1 + 300
		if (B-float64(b) == 468 || float64(b)-B == 468 ) { //B和b的关系 缺少  向浮点型转型 保留精度
			a = (b - 3) / 10
			fmt.Print(a, "\n")
			fmt.Println(B, b)
		}
	}
	c := 99
	//d := (c - 3) % 10
	//e := float64(d) * 0.1 //我想要e是浮点型 但是让的前身是整形，
	c1:=c*1.0 //可以
	//c2:=float64(c)*0.1 // 不可以 整型变量不可能不申明就变成浮点
	fmt.Println(c1)
//	fmt.Println(c2)

}
