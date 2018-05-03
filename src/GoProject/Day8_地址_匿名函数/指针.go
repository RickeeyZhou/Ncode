package main

import "fmt"

func stillTest(v int){
	v=v+10
}
func main() {
	i := 100
	fmt.Println("i", i)
	stillTest(i)
	fmt.Println("after i",i)
	anotherStillTest(&i)
	fmt.Println("after i",i)
	x:=1000
	fmt.Println("x",x)
	addressStillTest(&x)
	fmt.Println("after x",x)
}
func anotherStillTest(v *int ){
	*v=*v+100
}
func addressStillTest(v *int ){
	x:=456
	v=&x
}
