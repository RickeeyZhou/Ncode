package main

import "fmt"

func main()  {
	/*
	逻辑运算符：
	&&，逻辑与，"并且"
		所谓逻辑与运算符。如果两个操作数都非零，则条件变为真
		如果两个操作数都为true，结果才为true。
		逻辑与的判定规则：
			所有的操作数从左向右运算，
			如果遇到一个false就直接返回结果为false，
			除非所有的操作数都为true，结果才为true。
		"全真为真，一假为假"
	||，逻辑或，"或者"
		所谓的逻辑或操作。如果任何两个操作数是非零，则条件变为真
		逻辑或的判定规则：
			所有的操作数从左向右运算，
			如果遇到一个true就直接返回结果为true，
			除非所有的操作数都为false，结果才为false。
		"一真为真，全假为假"
	!，逻辑非，"就是取反"
		true，--->false
		false,--->true
	 */
	 b1 := false
	 b2 := false
	 b3 := true
	 res1 := b1 && b2 && b3 // 逻辑与
	 fmt.Println(res1)
	 res2 := b1 || b2 || b3 // 逻辑或
	 fmt.Println(res2)
	 res3 := !b1 // 逻辑非-->取反
	 fmt.Println(res3)

	 a := 3
	 b := 4
	 c := true

	 res4 := a < 4 && b / 2 >1 && c
	 //       T   		T      T
	 res5 := a * 2 < b && b % 2 != 0 && !c
	 //         F			F			F
	 res6 := (a + b)/b >= 0 || !c || b % 2 == 0
	//        T				F			T
	 res7 := !(a * 2 < b)
	 //       !F
	 fmt.Println(res4) // true
	 fmt.Println(res5) // false
	 fmt.Println(res6) // true
	 fmt.Println(res7) // true



}
