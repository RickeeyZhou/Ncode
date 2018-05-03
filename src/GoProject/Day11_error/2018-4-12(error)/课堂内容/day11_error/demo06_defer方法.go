package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p person) printInfo()  {
	fmt.Printf("姓名：%s，年龄：%d\n", p.name,p.age)
}

func main()  {
	p1 := person{"王二狗",30}
	p2 := person{"rose", 28}
	defer p1.printInfo()
	defer p2.printInfo()

	fmt.Println("over...")

	defer func() {
		fmt.Println("我是延迟的函数。。")
	}()
}

