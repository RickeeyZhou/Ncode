package main

import (
	"fmt"
)

func main() {
	var a, b int
	var r float32

	fmt.Scanln(&a,&b)
	fmt.Println( bijiao(a,b))
	fmt.Scanln(&r)
	s,l:=yuan(r)
	fmt.Println("s=",s,"l=",l)
}
func bijiao(a,b int )int{
	if a>=b{
	}else{
		a=b
	}
	return a
}
func yuan(r float32)( s float32,l float32){
	s1:=3.14*r*r
	l2:=6.28*r
	fmt.Printf("%p,%v,%p,%d\n",s,s,l,l)
	fmt.Printf("%p,%v,%p,%d2\n",s1,s1,l2,l2)
	return  l,s  //
}
func test ( a int )int{
	return 7
}
