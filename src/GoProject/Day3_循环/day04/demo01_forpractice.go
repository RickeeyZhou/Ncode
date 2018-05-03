package main

import "fmt"

func main() {
	/*

17.使用for循环，计算5的阶乘：5!=5*4*3*2*1

18.使用for循环，打印58-23数字

19.使用for循环，打印'A'到'Z'字母
		‘A’-->int32，数值
			65，'A'

20.使用for循环，打印1-100内的数字，每行打印10个。
	 */
	sum := 1
	for i := 5; i > 0; i-- {
		sum *= i
	}
	fmt.Println(sum)

	for i := 58; i >= 23; i-- {
		fmt.Println(i)
	}

	for i := 'A'; i <= 'Z'; i++ {
		fmt.Printf("%q\t", i)
	}
	fmt.Println()
	for i := 1; i <= 100; i++ {
		fmt.Print(i, "\t")
		if i%10 == 0 {
			fmt.Println() // print+line,直接换行
		}
	}
}
