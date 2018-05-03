package main

import "fmt"

type Address struct{
	city ,state string
}
type Person struct {
	name string
	age int
	address Address
}

func main(){
	var fancy Person
	fancy.name="付子龙"
	fancy.age=24
	fancy.address=Address{
		city: "New York",
		state:"????",
	}
	fmt.Println("Name：",fancy.name )
	fmt.Println("Age：",fancy.age)
	fmt.Println("Cityv：",fancy.address.city)
	fmt.Println("Cityv：",fancy.address.state)

}