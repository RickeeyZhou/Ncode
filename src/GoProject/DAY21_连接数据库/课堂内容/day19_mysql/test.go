package main

import (
	"fmt"
)

var s2 = []int{123, 51, 23, 43, 2, 5}

func main() {
	fmt.Printf("%p\n",s2) //0x52e600
	Section(s2)
	fmt.Printf("%p\n", s2) //0x52e600

	fmt.Println("截取后")

	fmt.Println(len(s2), s2) //6 [123 51 23 43 2 5]
}
func Section(ss []int) {
	fmt.Printf("%p\n",ss) //0x52e600
	//ss = append(ss[:len(ss)-2])
	ss = ss[:len(ss)-2]
	fmt.Printf("%p\n",ss) //0x52e600
	fmt.Println(cap(ss),len(ss), ss) //4 [123 51 23 43]
	//fmt.Printf("%p", ss)
	fmt.Println(cap(ss),len(s2),s2) //6 [123 51 23 43 2 5]

}