package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP: for a < 20 { // 10
	if a == 15 {
		/* 跳过迭代 */
		a = a + 1 // 16
		goto LOOP
	}
	fmt.Printf("a的值为 : %d\n", a) // 10, 11, 12 , 13,14,16,17,18,19
	a++//11 , 12 , 13 , 14,15
}
}