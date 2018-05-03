package main

import "fmt"

func main() {
	var age int
	fmt.Println("inter your age")
	fmt.Scanf("%d", &age)
	if age > 22 {
		fmt.Println("you have get the marred age")
	}

	fmt.Scanln(&age)
	if age < 0 {
		age = -1 * age
	}
	fmt.Println("绝对值是", age)
	var grade int
	fmt.Scanln(&grade)
	if grade > 60 {
		fmt.Println("Pass")
	}
}