package main

import "fmt"

type person2 struct {
	name string
	age int
}

func main()  {
	/*
	深拷贝：
		值类型：拷贝数据
	浅拷贝：
		拷贝地址-->指针

	结构体：
		值类型：
			p1 := person{.....}
		浅拷贝：
			指针：
				p2:=new(person)
				p3:= &p1
	 */
	 p1 := person2{"王二狗",30}
	 fmt.Printf("%p\n", &p1) //0xc0420023e0
	 p2 := p1 //值类型，深拷贝
	 fmt.Printf("%p\n",&p2)
	 p2.name="王二狗狗"
	 fmt.Println(p1)
	 fmt.Println(p2)

	 //定义一个指针,浅拷贝
	 p3:= &p1 // 将p1的地址给p3
	 fmt.Printf("%T\n", p3) //*main.person2
	p3.name ="李小花"
	fmt.Println(p1)
	fmt.Println(p3)
	fmt.Println(*p3)

	p4 := new(person2) // 指针
	fmt.Printf("%T\n",p4) //*main.person2
	p4.name="rose"
	p4.age = 30
	fmt.Println(*p4) //
	fmt.Printf("%p,%p\n", &p4, p4) //0xc042062020,0xc04203e4a0
	p6:=p4 //
	p6.name = "jack"
	fmt.Println(p4)
	fmt.Println(p6)

}
