package main

import "fmt"

func main()  {
	/*
	for 表达式1；表达式2；表达式3{
		循环体
	}

	 */
	 //1.省略表达式1
	i := 1
	 for ;i<5;i++{
	 	fmt.Println(i)//1,2,3,4
	 }
	 fmt.Println(i) // 5
	// 2省略表达2，相当于条件永远为true
	// for j:=0;;j++{
	// 	fmt.Println(j)
	// }

	// 3省略表达式3
	for j:=0;j<10;{
		fmt.Println(j)//0 0 0....
		j++//1
	}
	// 4.3个表达式都省略
	//for {//?for true
	//	fmt.Println(i)
	//	i++
	//}
	// 5. 省略表达式1和3
	k :=0
	for k < 5{ // for ; k<5;
		fmt.Println(k) // 0 1 2 3 4
		k++ // 1 2 3 4 5
	}
}
