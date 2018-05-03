package main

import "fmt"

type student struct {
	Name string
}
func zhoujielun (v interface{})string{
	switch msg := v.(type){
	case *student:
		 return msg.Name
	case student:
		fmt.Println(msg)
		return msg.Name
	case string :
		fmt.Println(msg)
		return msg
	}
	return "error"
}
// msg 是一个接口类型,不能调用结构体属性的
func main(){
	s:="jilao"
	dengchao:=student{"dengchao"}
fmt.Println(zhoujielun(s))
fmt.Println(zhoujielun(dengchao))
fmt.Println(zhoujielun(&dengchao))
}