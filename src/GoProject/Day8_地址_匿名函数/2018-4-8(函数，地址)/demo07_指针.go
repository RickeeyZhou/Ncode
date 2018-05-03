package main

import "fmt"

func main()  {
	/*
	指针：存储另一个变量的内存地址的变量。
	指针的类型，
	指针存储的数据的类型
	指针的地址
	指针存储的变量的地址
	 */
	 a:=100 //定义了一个变量，变量名：a，变量值：100，变量类型int
	 fmt.Printf("变量的数值：%d，变量的类型：%T\n",a,a)
	 //&,获取变量的内存地址
	 fmt.Printf("变量a的内存地址：%p\n",&a) //0xc04200e098

	 // 修改变量的值
	 a = 200//修改a的数值
	 fmt.Println(a)
	 fmt.Printf("%p\n",&a)

	 //创建一个指针
	 var p1 *int
	 fmt.Println(p1) //<nil> 空 null
	 fmt.Printf("%T\n",p1) //*int
	 fmt.Printf("p1自己的地址：%p\n",&p1)
	 p1 = &a //取a的地址，赋值给p1
	 fmt.Println(p1)
	 fmt.Println("获取数值：",*p1)

	 // 指针的的指针

	 var p2 **int
	 fmt.Println(p2) //nil
	 fmt.Printf("%T\n",p2) //**int
	 p2 = &p1 // int-->*int
	 fmt.Printf("p2中存储的数据，就是p1的地址：%p\n",p2)//p2的值，就是p1的地址
	 fmt.Printf("p2自己的地址：%p\n",&p2)//p2的地址：

	 fmt.Println(*p1) //*,获取p1指针中存储的内存地址对应的数据。 200
	 fmt.Println(*p2) // *p2取值的a的地址
	 fmt.Println(**p2)




}
