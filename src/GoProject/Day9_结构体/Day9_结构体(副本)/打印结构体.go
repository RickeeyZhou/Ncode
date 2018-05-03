package main

import "fmt"

type Employee struct {

	id string
	name string
	salary int

}
func main(){
	var s1 Employee
	s1=Employee{"1304030304","Rickeey",1000000}
	s2:=Employee{name:"Jerome",id:"f**k off",salary:99999}
	s3:=Employee{}
	s3.salary=10000
	s3.name="Tank"
	s3.id="shit"
	fmt.Println(s1,"\t",s2,"\t",s3)
	s4:=new(Employee)
	s4.id="dammmmm"
	s4.name="joy"
	s4.salary=1231231312
	fmt.Println(*s4)
}
