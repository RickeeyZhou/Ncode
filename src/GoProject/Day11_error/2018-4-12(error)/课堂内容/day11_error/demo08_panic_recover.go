package main

import "fmt"

func main()  {
	/*
	go中的异常处理：
		用defer panic recover
		defer ：延迟执行
		panic："恐慌","害怕"
			打断程序的正常执行。
		recover："恢复"
	 */

	 testA()
	 testB()
	 testC()
	 fmt.Println("main...over...")
}
func testA()  {
	fmt.Println("函数A（）。。。。")
}

func testB()  {
	// recover处理函数中的panic，
	// 恢复panic所中断的程序
	defer func() {
		if msg:=recover();msg !=nil{ //
			fmt.Println(msg)
		}
	}()

	for i:=1;i<=10;i++{
		if i == 5{
			panic("testB函数，panic..") // 恐慌-->程序到此中断
		}
		fmt.Println("函数B,",i) // 1,2,3,4
	}
}

func testC()  {
	defer func() {
		recover() // 恢复程序：
		/*
			程序中没有panic，recover(),接收到的nil，
			程序中有panic，接收到panic的信息。。
		 */
	}()
	panic("函数c，panic。。。")
	fmt.Println("函数C()....")
}
