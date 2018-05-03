package main

import "fmt"

func main() {
	fmt.Println(`Print 1~100 number ,and 10/line `)
	printnum()
	printabc()
	var a int
	fmt.Scanln(&a)
	fmt.Println(factorial(a))
	fmt.Println("................................................")
	addmin()
	var year int
	fmt.Println("year ")
	fmt.Scanln(&year)
	fmt.Println(runian(year))
}
func printnum() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i, "\t")
		if (i)%10 == 0 {
			fmt.Println()
		}
	}
}
func printabc() {
	num := 1
	var i int
	for i = 'A'; i <= 'z'; i++ {
		fmt.Printf("%c \t ", i)
		num++
		if num%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Printf("\n%T,%q,%c ,%v", i, i, i, i)
}
func factorial(a int) (sum int) {
	sum = 1
	for i := 1; i <= a; i++ {
		sum *= i
	}
	return
}
func addmin() {
	var a, b int
	var c string
	fmt.Scan(&a)
	fmt.Scan(&c)
	fmt.Scan(&b)
	switch c {
	case "*":
		fmt.Printf("%d*%d=%d \n", a, b, a*b)
	case "/":
		fmt.Printf("%d/%d=%d \n", a, b, a/b)
	case "+":
		fmt.Printf("%d+%d=%d \n", a, b, a+b)
	case "-":
		fmt.Printf("%d-%d=%d \n", a, b, a-b)
	}
	return
}
func runian(a int) bool {
	if a%4 == 0 && a%100 != 0 || a%400 == 0 {
		return true
	} else {
		return false
	}
}
