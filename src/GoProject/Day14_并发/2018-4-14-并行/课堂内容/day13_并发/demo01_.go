package main

import "fmt"
/*
go执行程序的时候，先创建了一个goroutine，叫主goroutine。执行main()

 */
func main() {
	fun2()
	fun3()
}
func fun1() {
	fmt.Println("我是fun1.。。。")
}
func fun2() {
	for i := 0; i < 10; i++ {
		fmt.Println("fun2..")
	}
	fun1()
}

func fun3()  {
	fmt.Println("我是fun3")
	fun1()
}
