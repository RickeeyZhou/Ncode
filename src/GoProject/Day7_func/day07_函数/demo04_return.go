package main

import "fmt"

func main()  {
	/*
	return：
		1.返回函数的结果给调用处
		2.结束函数的执行

	注意事项：
	1.如果函数声明上有返回值，那么函数中必须有return返回结果。无论函数中有分支，循环，保证一定有结果可以返回。
	2.如果函数没有声明返回值，那么函数中也可以使用return关键字，用于强制结束函数
	3.return后的数据，要保持和声明的返回值类型，数量，顺序一致
	 */
	 fun1()
	 fun2()
	 fun3()
	 fmt.Println(fun4(30))
}

func fun1(){
	age := -28
	if age < 0{
		return // 强制结束函数
	}
	fmt.Println(age)
}
func fun2()  {
	for i:=1;i<=10;i++{
		if i == 5{
			//break // 结束循环
			return // 结束函数
		}
		fmt.Println(i)//1 2 3 4
	}
	fmt.Println("over。。。")
}

func fun3(){
	fmt.Println("hello")
	return
	//fmt.Println("world")
}

func fun4(age int) int  {
	if age < 0{
		return 0
	}

	return age
}
