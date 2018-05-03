package main

import "fmt"

func main()  {
	/*
	函数：
		就是具有一定功能的一段代码。封装好了的一段代码。可以被多次调用执行。
	作用：
		1.可以避免重复的代码
		2.增强了程序的扩展性

	使用：
		A：函数的定义，声明
		B：函数的调用。而且可以多次调用
	语法：
		func funcName(parametername type1, parametername type2) (output1 type1, output2 type2) {
		//这里是处理逻辑代码
		//返回多个值
		return value1, value2
	}
		func：关键字
		funcName：函数名字，就是一个标识符
			首字母大写：公共的，首字母小写：私有的
		参数列表：函数执行时，需要外部传入数据
		返回值：函数执行后，返回给调用处的结果
		(),函数的标志

	注意事项：
		1.函数必须先声明定义，然后再进行调用
		2.函数调用：函数名()
		3.函数名不能冲突
		4.main(),是程序的入口，自动调用执行。
	 */
	 //1.求1-10的和
	//进行函数的调用
	getSum() // 函数的调用处
	 fmt.Println("HelloWorld..")
	 //求1-10的和
	getSum()

	fmt.Println("王二狗")
	//1-10的和
	getSum()
}

/*
定义一个函数：用于求1-10的和
 */

 func getSum(){
	 sum := 0
	 for i:=1;i<=10;i++{
		 sum += i
	 }
	 fmt.Println("1-10的和是：",sum)
 }

 /*
 1.定义一个函数，用于求5的阶乘。然后进行调用
2.定义一个函数，打印三角形，然后进行调用。
  */
