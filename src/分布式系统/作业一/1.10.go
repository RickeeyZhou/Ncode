package main

import "fmt"

type People1 interface{
	Speaking(string)string
}
type Student1 struct {

}
func (stu *Student1)Speaking (think string )(talk string ){
	if think =="bitch"{
		talk="you are a goood boy"

	}else {
		talk ="hi"
	}
	return
}
func main(){
	var peo Student1
	peo=Student1{} //类型不同
	think :="bitch"
	fmt.Println(peo.Speaking(think))
}
// 讨论接口的运用,不能直接将接口对等 结构体,
//只能是某个结构体对象实现了接口