package main

import "fmt"

func main() {
	var n int
	var count int
	var sum int
	fmt.Scanln(&n)
	arr := make([]int, n, 100)
	for k := 0; k < n; k++ {
		arr[k] = 1
	}
 	i:=0
	//	for j := 1; j < n; j++ {

	for sum < n-1 {           //固定循环 n-1次，每一循环，派出一个人，
		if arr[i] == 1 {	//用下标自加去遍历整个数组
			count++
			if count == 3 {
				count = 0
				arr[i] = 0
				sum += 1
			}
		}
		i++
		if i == n {
			i = 0
		}
		//if sum == n-1 {
		//	break
		//}
	}
	//}
	fmt.Println("flag")
	fmt.Println(n)
	fmt.Println(arr)
	for j := 0; j < n; j++ {
		if arr[j] == 1 {
			fmt.Println(j)
		}
	}

}
