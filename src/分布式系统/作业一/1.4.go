package main

import "fmt"

type People struct {
	// 这是一个空结构体
}
func (p *People)ShwoA(){  ///showa的方法
	fmt.Println("ShowA")
	p.ShowB()
}
func (p *People)ShowB(){
	fmt.Println("ShowB")
}
type Teacher struct{
	People    // 继承了People
}
func (t *Teacher) ShowB(){
	fmt.Println("teacher showB")
}
func main(){
	t:=Teacher{}
	t.ShwoA()
}
/*
输出结果: SHowA
		ShowB
 */