package main

import "fmt"

type Employee struct{
	string
	int
}

func main(){
e2:=Employee{int:30,string:"lee"}
fmt.Println(e2)
changeName(e2)
changeAge(&e2)
//////////////////////////////////////////////////////////
b2:=struct{
	int
	string
}{
	2,
	"232",
}
fmt.Println(b2)
}
func changeName(p Employee ){
	p.string="fuck"
	fmt.Println(p)
}
func changeAge(p *Employee){
	(*p).int=26
	fmt.Println(p)
	p.int=33
	fmt.Println(*p)
}
