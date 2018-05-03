package main

import "fmt"

type person1 struct {
	name string
	age int
	sex string
}
func main(){
	var p1 person1
	fmt.Println(p1)

	p1.name="王二狗"
	p1.age=30
	p1.sex="male "
	fmt.Println(p1)
	fmt.Println("name:",p1.name,"age",p1.age,"sex:",p1.sex)
	fmt.Printf("%T,%p\n,..%T,%p",p1,&p1,p1.sex,&p1.sex)

	p2:=person1{"fancy",18,"female"}
	fmt.Println(p2)
	fmt.Printf("%T,%p\n",p2,&p2)

	p3:=person1{age:30,sex:"male ",name:"fuck"}
	var p4 person1
	p4=person1{
		name:"kk",
		sex:"female",
		age:23,
	}
	fmt.Printf("%T,%p\n",p3,&p3)
	fmt.Println(p4)

	p5:=new(person1)
	fmt.Printf("%T\n",p5)
	(*p5).name="neibor"
	p5.age=30
	p5.sex="boy"
	fmt.Println(*p5)
}
