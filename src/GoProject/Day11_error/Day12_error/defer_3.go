package main

import "fmt"

func main(){
	p1:=person{"er_Dog",40}
	p:=person{"rose",28}
	defer p1.printInfo()
	defer p.printInfo()
	fmt.Println("over....")
	defer func (){
		fmt.Printf(" i am defer function...")
	}()
}
type person struct {
	name string
	age int
}
func (p person )printInfo(){
	fmt.Printf("name:%s,age:%d\n",p.name,p.age)
}