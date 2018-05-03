package 接口

import "fmt"

type Person struct {
	Name string
	Age int
}
func(a Person) PrintInfo(){
	fmt.Println(a.Name,a.Age)
}
type Student struct{
	Person
	Grade int
}
type Employee struct {
	Person
	Salary int
}
