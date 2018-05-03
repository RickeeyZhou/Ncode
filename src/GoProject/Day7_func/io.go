package main

import"fmt"
var V1 int32
var V2 float32
var V3 bool
var V4 string

func main(){
	fmt.Println("keep moving!")
	fmt.Print("it 's a sunny day ~")
	fmt.Print(" and i feel better \n")
	fmt.Println("the thing we have do must be done on time ")
	fmt.Print(".....")
	V1= 100
	V2=10.21
	V3=true
	V4="flase"
	fmt.Printf("%d,%T\n", V1,V1)
	fmt.Printf("%e,%T\n" ,V2,V2)
	fmt.Printf("%t,%T\n" ,V3,V3)
	fmt.Printf("%v,%T\n" ,V4,V4)
}
// define five reasonable variable
