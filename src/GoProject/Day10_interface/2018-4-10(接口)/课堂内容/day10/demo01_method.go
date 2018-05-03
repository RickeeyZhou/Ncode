package main

import "fmt"

//1.定义一个结构体
type Person struct {
	name string
	age int
	sex string
}

type Dog struct {
	sort string
	color string
}

// 2.定义一个方法
func (p Person) printInfo(){ // p-->p1
	//可以直接操作Person类的这个对象，就是p
	fmt.Printf("姓名：%s,年龄：%d,性别：%s\n",p.name, p.age,p.sex)
}

func (d Dog)printInfo()  {
	fmt.Printf("品种：%s,毛色：%s\n", d.sort,d.color)
}

func printInfo2(p Person)  {
	fmt.Printf("姓名：%s,年龄：%d,性别：%s\n",p.name, p.age,p.sex)
}



func fun1(a, b int)int{
	return a + b
}

func main()  {
	/*
	方法：method---->类的行为功能。
		接受者调用--->
		func (接受者) 方法名 (参数列表)(返回值列表){
			//...
		}
	函数：function
		可以直接调用的
		func 函数名(参数列表)(返回值列表){

		}

	方法和函数的区别：
		函数是定义功能，一段代码
		方法的某个类别的行为功能

		方法由于接受者不同，方法名可以一直，但是函数不能。go中没有函数重载
	 */

	 //1.创建一个对象
	 p1 := Person{"王二狗",28,"男"}
	 p1.printInfo()

	 p2 := Person{"李小花",18,"女"}
	 p2.printInfo()

	d1 := Dog{"中华田园犬","土黄色"}
	d1.printInfo()
	d2:=Dog{"哈士奇","灰白色"}
	d2.printInfo()



	 printInfo2(p1) // 函数调用
	 printInfo2(p2) // 函数调用
}
/*
定义一个car的结构体，
	字段有名字，速度。
定义一个方法，run(),打印该车的速度
 */
