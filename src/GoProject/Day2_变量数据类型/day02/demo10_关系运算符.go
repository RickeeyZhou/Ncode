package main

import "fmt"

func main(){
	/*
	关系运算符：运算结果是bool
	>,<,>=,<=, ==, !=
	==,判断两个数值是否相等
	!=,判断两个数值是否不等
	 */
	 a := 8
	 b := 5
	 fmt.Println(a < b)// 8 < 5,false
	 fmt.Println(a == b) // false
	 fmt.Println(a != b) // true

	 c := true
	 d := false
	 //fmt.Println(c < d) // bool类型不能比较大小
	 fmt.Println(c == d) // false

}
