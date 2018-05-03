package main

import "fmt"

func main() {
	//1-5 的和
	sum := getsum(5)
	fmt.Println(sum)

	fmt.Println(jie(5))

	fmt.Println(fibo(12))
}
func getsum(n int) int {
	fmt.Println("*")
	if n == 1 {
		return 1
	}
	return getsum(n-1) + n

}

/*
练习1：递归算法，求5的阶乘，
练习2：递归算法，求fibonacci的第12项数值
 */

func jie(n int) int {
	if n == 1 {
		return 1
	}
	return jie(n-1) * n
	/*
	return jie(4)*5
	return jie(3)*4
	rerun jie(2)*3
	return jie(1)*2
	return 1
	 */
}

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
