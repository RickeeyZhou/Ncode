package main

import "fmt"

func main(){
	const a="b"
	fmt.Print(a)


	var s4  =`A`
	fmt.Printf("s4的数据类型 %T他的值是%q %s \n", s4,s4,s4)
	a1 := 100
	b := 322222222222222222333333333333333333333333.1433
	c := "hello"
	d := `王二狗住在隔壁`
	e := true
	f := 'A'
	fmt.Printf("%d,%f,%s,%s,%v,%q",a1,b,c,d,e,f)
}
