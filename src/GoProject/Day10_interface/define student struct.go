package main

import "fmt"

func main(){
	me:=studenthk{name:"Rickeey Zhou",age:23,sex:"male",mygrade:grade{E_grade:99,M_grade:98,C_grade:85}}
	gradesum,gradeaver:=me.mygrade.grade_de()
	fmt.Println(gradesum,gradeaver)
	fmt.Println(me.name,me.age,me.sex)
}
type studenthk struct {
	name string
	age int
	sex string
	mygrade grade
}
type grade struct {
	E_grade int
	C_grade int
	M_grade int
}
func ( a grade )grade_de()(sum int, average float64){
	sum = a.C_grade+a.E_grade+a.M_grade
	average=float64(sum)/3
	return
}
