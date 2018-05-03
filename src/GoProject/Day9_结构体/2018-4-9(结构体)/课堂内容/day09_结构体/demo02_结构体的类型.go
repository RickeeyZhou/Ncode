package main

import "fmt"

/*
定义dog类
 */
 type dog struct {
 	name string
 	color string
 	age int
 }

func main()  {
	/*
	结构体和数值一样，是值类型
	 */
	d1 := dog{"啸天","黑色",3}
	d2 := d1
	d2.name="小二黑"
	fmt.Println(d1) //{啸天 黑色 3}
	fmt.Println(d2) //{小二黑 黑色 3}

	d3 :=new(dog)
	d3.name="二黄"
	d3.color="黄色"
	d3.age = 2

	d4 := d3
	fmt.Printf("%T,%T\n",d3,d4)//*main.dog,*main.dog
	d4.name="大黄"
	fmt.Println(d3) //&{大黄 黄色 2}
	fmt.Println(d4) //&{大黄 黄色 2}

	d5 :=&d1 //定义了一个指针
	fmt.Printf("%T\n",d5)//*main.dog
	d5.name="黑狗"
	fmt.Println(d1) //{黑狗 黑色 3}
}
/*
练习1：创建一个结构体：学生
练习2：创建结构体对象，s1，s2，通过代码值类型，
 */
