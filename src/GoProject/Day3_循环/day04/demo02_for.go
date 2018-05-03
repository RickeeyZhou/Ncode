package main

import "fmt"

func main() {
	/*
练习1：打印1-100内，所有的偶数
	练习3：求1-100内，奇数的和
练习2：打印1-100内，能被3整除，不能被5整除的数，每行打印5个
统计打印数字的个数。
	 */
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	// 2, 4, 6, 8, 10....
	for i := 2;i <= 100;i+=2{
		fmt.Println(i)
	}

	sum := 0
	for i:=1;i<=100;i+=2{
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("--------------")
	count := 0//用于计数器
	for i:=1;i<=100;i++{
		if i % 3 == 0 && i % 5 != 0{
			count++ // 5, 10, 15...
			fmt.Print(i,"\t")
			if count % 5 == 0{
				//fmt.Println()
				break //彻底的结束循环
			}
			//count++
		}
	}
	fmt.Println()
	fmt.Println("count-->", count)


}
