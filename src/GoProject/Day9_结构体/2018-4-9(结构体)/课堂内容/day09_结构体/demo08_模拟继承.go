package main

import "fmt"

type Person5 struct {
	name string
	age int
}
//在结构体中属于匿名结构体的字段称为提升字段
type Student struct {
	//name string
	//age  int
	Person5 // 模拟继承：结构体的嵌套
	school string
}


func main()  {
	/*
	继承：两个类的继承关系。子类可以使用父类的功能
	对于go语言，没有真正的继承。靠结构体的嵌套模拟继承功能的。
	结构体的嵌套：
		聚合关系：一个类作为另一个类的属性。has - a
			type A struct{}
			type B struct{
				a A
			}
		继承关系：一个子类继承一个父类。is - a
			type C struct{}
			type D struct{
				C
			}

	 */
	 p1 := Person5{"王二狗",30}
	 fmt.Println(p1.name,p1.age)

	 s1 := Student{}
	 s1.name ="李小花" // 提升字段
	 s1.age=30 // 提升字段
	 s1.school="清华小班"
	 fmt.Println(s1.name,s1.age,s1.school)

	 s2:=Student{Person5{"如梦",20},"清华大班"}
	 fmt.Println(s2)
	 s3:=Student{Person5:Person5{"如歌",23},school:"清华中班"}
	 fmt.Println(s3)

	 /*
	 练习：
	 父类：Animal
	 子类：
	 	Cat，Dog
	 创建子类对象，为属性赋值，并打印结果。
	  */
}
