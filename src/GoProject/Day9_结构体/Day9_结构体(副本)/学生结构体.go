package main

import "fmt"

type student struct{
	name string
	age int
	sex string
	Grade grade
}
type grade struct{
	math int
	chinese int
	english int
}
func printInfo( a student){
	var sum, x int
	sum=a.Grade.chinese+a.Grade.english+a.Grade.math
	x=(a.Grade.chinese+a.Grade.english+a.Grade.math)/3
	fmt.Println(a.name,a.sex,a.age)
	fmt.Println("Sum:",sum,"average:",x )
}
func main(){
	fuzilong:= student{
		name:"fuzilong",age:24,sex:"male",Grade:grade{99,88,22},
	}

	printInfo(fuzilong )
}
