package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
	Sex string
}

func (p Person)Say(msg string)  {
	fmt.Println("hello，",msg)
}
func (p Person)PrintInfo()  {
	fmt.Printf("姓名：%s,年龄：%d，性别：%s\n",p.Name,p.Age,p.Sex)
}

func main()  {
	//1.创建一个struct的对象
	p1 := Person{"王二狗",30,"男"}
	//1.第一部分：通过反射操作结构体中的字段：获取字段的类型，名称，数值。。
	//通过reflect包下的TypeOf()，获取p1的类型

	t1:=reflect.TypeOf(p1)
	fmt.Println(t1) //main.Person
	fmt.Println(t1.Name()) //Person
	fmt.Println(t1.Kind()) //struct
	//通过reflect包下ValueOf(),获取p1的值
	v1:=reflect.ValueOf(p1) //{王二狗 30 男}
	fmt.Println(v1)

	//详细操作：判断t1的种类是结构体类型
	if t1.Kind() == reflect.Struct{
		//获取字段的数量
		count:=t1.NumField() //3个字段
		fmt.Println(count)

		for i:=0;i<count;i++{
			filed:=t1.Field(i)
			val:=v1.Field(i).Interface()//根据v1获取对应字段的数值，--->空接口类型
			//fmt.Println(filed)
			fmt.Println("字段名称：",filed.Name,"字段类型：",filed.Type,"字段数值：",val)
		}
	}


	//2.通过反射，操作方法
	fmt.Println(t1.NumMethod())
	for i:=0;i<t1.NumMethod();i++{
		method:=t1.Method(i)
		fmt.Println(method.Name,method.Type)
	}

	//执行方法
	m1:=v1.MethodByName("PrintInfo")
	m1.Call(nil)//执行m1对应的PrintInfo()

	m2:=v1.MethodByName("Say")
	m2.Call([]reflect.Value{reflect.ValueOf("你好啊")}) // string-->Value

}
