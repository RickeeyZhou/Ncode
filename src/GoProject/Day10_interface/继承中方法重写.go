package main

import "fmt"


type human struct {
	name string
	age int
}
type book struct {
	bookname string
}
type student struct {
	school string
	human
	mybook book
}

func (people human )eat(){
	fmt.Println("eating ...")
}
func(people human )printInf (){
	fmt.Println(people.name,people.age)
}
func (child student) eat(){
	fmt.Println("i have reride the function")
}
func (child student)study(){
	fmt.Println("i'm studying ")
}
func main(){
	fancy:=student{school:"Newyork university",human:human{"fancy",24}}
	fancy=student{mybook:book{"porn.."}}  //get the value again
	fancy.eat()
	fancy.study()
	fmt.Println(fancy)

}