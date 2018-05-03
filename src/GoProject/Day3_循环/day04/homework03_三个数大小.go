package main //

import "fmt"//

func main(){
	/*
	给定3个整数，从小到大输出。
	 */
	 a := 2
	 b := 3
	 c := 1
	 //step1：比较a和b
	 if a > b{
	 	a, b = b, a
	 }
	// step2：比较a 和c
	if a > c{
		a, c = c, a
	}
	//step3：比较b和c
	if b > c{
		b, c = c, b
	}
	fmt.Println(a," <= ", b," <= ", c)
}
