package main

import "fmt"
/*
函数的参数：
	形式参数列表：
		函数在执行过程中，需要由外部传入数据。
	实际参数列表：
		调用函数时，传入的具体的数据，叫做实参

调用：
	A：函数名必须一致
	B：实参严格匹配形参：一一对应
		类型，个数，顺序


func 函数名(参数名 参数类型,参数名 参数类型){

}

参数列表： 参数名 参数类型
	如果多个参数类型一致，可以省略成一个。
 */
func main() {
//1.1-20的和
	num:=20
	getSum2( num) // 20
//2.1-30的和
	getSum2(30) // 30
//3.1-100的和
	getSum2(100)//100

	getAdd1(20,10)
}
// 求1-n的和
func getSum2(n int) {//形参
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1-%d的和：%d\n",n, sum)
}

func getAdd1(num1, num2 int){//num1 int, num2 int
	fmt.Println("num1:",num1)
	fmt.Println("num2:",num2)
	sum := 0
	sum = num1 + num2
	fmt.Printf("%d + %d = %d\n",num1,num2,sum)
}

func test1(a, b int, c string){ // 10, 20, ""

}

func test2(c string, a, b int)  {// "",10,20

}

/*
1.定义一个方法：求n的阶乘。调用的时候，由键盘输入。
2.定义一个方法：求2个数比较大小。2个数由参数传入。。
 */