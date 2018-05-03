package employee

import "fmt"

type Employee struct {
	Name string
	Age int
	Salary float64
}

func (e Employee) PrintInfo()  {
	fmt.Printf("姓名：%s，年龄：%d,工资：%.2f\n",e.Name,e.Age,e.Salary)
}



