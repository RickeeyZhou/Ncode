package main

import "fmt"

func main()  {
	// iota，特殊的常量值，每当定义了const，iota的值置为0，
	// 每当定义一个常量，累加1。遇到const时，重新置为0
	const  (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	fmt.Println(a, b, c)

	// 简写成
	const (
		d = iota
		e
		f
	)
	fmt.Println(d, e, f)

}
