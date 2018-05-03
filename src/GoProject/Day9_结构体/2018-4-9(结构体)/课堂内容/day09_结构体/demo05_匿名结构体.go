package main

import "fmt"

type Employee struct {
	string //匿名字段
	int
	//sex string
}
// 匿名字段：没有名字的字段，但是有类型。相当于该类型就是字段名
//注意点：一个结构体中，匿名字段的类型是唯一的。

func main()  {
	/*
	匿名结构体和匿名字段
		匿名结构体：没有名字的结构体

	创建结构体对象：
		对象名:=结构体名{field:value,field:value...}

	匿名字段：
	 */

	 //s1:=a{}
	 s2:=struct {
		name string
		age int
	}{
		name :"王二狗",
		age :30,
	 }
	 fmt.Println(s2)

	 fmt.Println("---------------")
	 //e1 := Employee{"王二狗",30}
	 e2 := Employee{int:30,string:"李小花"}
	 fmt.Println(e2)
	 fmt.Println(e2.string)
}
