package main

import "fmt"

func main() {
	fmt.Printf("this is a caculate Programming")
	fmt.Printf("plz import two number")
	var a,b int
	var c string
	fmt.Scanf("%d ,%d",&a,&b)
	fmt.Scanln(&c)
	switch c{
	case `+`:
		fmt.Println(a+b)
	case `-`:
		fmt.Println(a-b)
	case `*`:
		fmt.Println(a*b)
	case `/`:
		fmt.Println(a/b)
	default:
		fmt.Println("非法运算")
	}

}
