package main

import "fmt"

type people1 struct{
	name string
}
func ( p *people1) String() string {
	return (fmt.Sprintf("print:%v",p.name))

}
func main(){
	p:=&people1{"神神道道"}
	fmt.Println(p.String())
}
// 题目利用了 Sprintf ( string, interface{}),接口接受任何值的特性
//理论上给任何值都不会报错.
// string 函数要求返回 string类型, 编译器会去检查springtf,是不是返回string
//但是如果是赋值 p,
//p是一个结构体对象的地址
// 还是不懂为什么错