package main

import (
	"fmt"
	"strconv"
)

func main(){

	s1:="200"
	fmt.Println(s1)

	i1,err:=strconv.ParseInt(s1,10,64)
	fmt.Println(i1,err)
	fmt.Printf("%T\n",i1)

	s2:="true"
	b1,err:=strconv.ParseBool(s2)
	fmt.Println(b1,err)

	num1:=100
	s3:=strconv.FormatUint(uint64(num1),10)
	s4:=strconv.FormatBool(false)
	fmt.Printf("%T,%s\n",s4,s4)

	i2,err:=strconv.Atoi(s1)
	fmt.Println(i2)
	s5:=strconv.Itoa(1100)
	fmt.Println(s5)
}