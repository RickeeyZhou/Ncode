package main

import "fmt"
func main() {
	var a int
	fmt.Println(a)
	var a1[2]int
	fmt.Printf("%p",a1)
	var a3[]int
	fmt.Println(a3)
	var a4 map[int]string
	fmt.Printf("%p",a4)
	var a5 []int

	fmt.Printf("%p",a5)
}