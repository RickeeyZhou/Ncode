package main

import "fmt"

func main()  {
	/*
	数据类型的转换
	强类型语言，运算时，需要统一类型。
	T(v)
		T:type
		v:value
	 */
	 a := 100 // int
	 b := 3.84 //float64
	 c := int(b) // float64--->int
	 sum := a + c
	 fmt.Println(sum)

	 d := float64(a)
	 fmt.Println(d)
}
