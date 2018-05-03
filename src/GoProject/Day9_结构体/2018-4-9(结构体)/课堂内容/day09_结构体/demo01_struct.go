package main

import "fmt"

/*
定义一个结构体
type 结构体名 struct{
	字段名 类型
}
 */
type person1 struct {
	name string //姓名
	age int // 年龄
	sex string // 性别
}

func main()  {
	//根据类实例化对象：根据结构体创建该结构体类型的变量
	//1.
	var p1 person1
	fmt.Println(p1) //{ 0 }
	// 对象.属性= 赋值
	p1.name = "王二狗"
	p1.age = 30
	p1.sex = "男"
	fmt.Println(p1)//{王二狗 30 男}
	fmt.Println("姓名：",p1.name,"年龄",p1.age,"性别：",p1.sex)
	fmt.Printf("%T,%p\n",p1, &p1) //main.person1,0xc042076060

	//2.
	p2 := person1{"李小花", 18,"女"} //注意赋值顺序
	fmt.Println(p2)
	fmt.Printf("%T,%p\n",p2,&p2) //main.person1,0xc042076120

	//3
	p3 := person1{age:30,sex:"男",name:"三胖"}
	p4 := person1{
		name:"rose",
		age:30,
		sex:"女",
	}
	fmt.Printf("%T,%p\n",p3,&p3) //main.person1,0xc0420781b0
	fmt.Println(p4) //{rose 30 女}
	//4.new()，内置函数，用于创建指针
	p5 := new(person1)
	fmt.Printf("%T\n",p5) //*main.person1
	(*p5).name="隔壁老王"
	p5.age = 30 // 简写
	p5.sex = "男"
	fmt.Println(p5) //&{隔壁老王 30 男}


}
