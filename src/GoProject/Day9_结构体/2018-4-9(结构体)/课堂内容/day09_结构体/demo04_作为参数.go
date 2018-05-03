package main

import "fmt"

type person3 struct {
	name string
	age int
}

func main()  {
	p1 := person3{"王二狗", 30}
	fmt.Println(p1)
	changeName(p1)
	fmt.Println(p1)

	changeName2(&p1)
	fmt.Println(p1) //{李小花 30}
}

func changeName(p person3)  { // p = p1,值传递
	p.name = "西哲"
	fmt.Println(p)
}

func changeName2(p *person3)  { // p,指针类型 p = &p1
	p.name ="李小花"
	fmt.Println(p,*p) //&{李小花 30}
}
