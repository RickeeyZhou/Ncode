package main

import (
	"fmt"
	"math"
)

//1.定义Animal的结构体
type Animal struct {
	name string
	age int
}
//2.定义Animal的方法
func (a Animal) eat()  {
	fmt.Println(a.name,",在吃东西。。。")
}
func (a Animal) printInfo()  { // Anmial的对象调用方法
	fmt.Printf("名字：%s，年龄：%d\n", a.name,a.age)
}

// 3.定义Cat和Dog继承
type Cat struct {
	Animal
	color string
}
// 因为eat(),在父类Animal已有，但是Cat重新实现了，所以叫重写：override
func (c Cat) eat(){
	fmt.Printf("小猫：%s,在吃鱼。。。\n",c.name)
}

type Dog2 struct {
	Animal
	sex string
}
// 子类新增的方法
func (d Dog2) lookDoor()  {
	fmt.Printf("%s,在看家。。\n",d.name)
}

func main()  {
	/*
	继承：结构体的嵌套
		子类可以直接访问父类的属性和方法
		子类可以新增自己的属性和方法
		子类可以重写父类的方法
	 */
	 a1 := Animal{"兔子",3}
	 a1.eat()
	 a1.printInfo()

	 c1 := Cat{Animal{"Tom",3},"蓝灰色"}
	 c1.eat()// 子类访问重写的方法。
	 c1.printInfo()// 子类直接访问父类的方法

	 d1:= Dog2{Animal{"啸天",2},"公狗"}
	 d1.eat() //子类对象，访问父类的方法
	 d1.printInfo() //

	 d1.lookDoor()

	 fmt.Println(math.Sqrt(8))//根号8,
	 fmt.Println(math.Sqrt(9))
}
/*
练习1：
父类：人类
	name，age
	eat(),printInfo()

子类：学生类
	school
	重写eat(),新增study()

练习2：
定义一个类：Point
	x，y，z轴
创建点的对象：
	两个点

定义函数，方法，求两点之间的距离


练习3：定义学生类
	姓名，年龄，性别，英语成绩，语文成绩，数学成绩
	方法：
		求总分()
		求平均分()
		打印学生信息()
 */
