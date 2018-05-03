package main

import (
	"reflect"
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

type Cat struct {
	Animal
	Color string
}

func main() {
	c1 := Cat{Animal{"猫咪", 1}, "白色"}

	t1 := reflect.TypeOf(c1)
	fmt.Println(t1.NumField())

	for i := 0; i < t1.NumField(); i++ {
		field := t1.Field(i)
		fmt.Println(field)
	}

	//获取匿名字段
	/*
	匿名字段，可以通过下标获取
	FieldByIndex（[]int{0}）-->Animal
	FieldByIndex（[]int{0,0}）-->Animal中的第一个字段：Name
	 */
	f1:=t1.FieldByIndex([]int{0,0})
	fmt.Println(f1.Name,f1.Type)//Name string
	f2:=t1.FieldByIndex([]int{0,1})
	fmt.Println(f2.Name,f2.Type)


	v1:=reflect.ValueOf(c1)

	name:=v1.FieldByIndex([]int{0,0})
	fmt.Println(name)
}
