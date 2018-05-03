package main

import "fmt"

type wheel struct{
	inside_radius float64
	outside_radius float64
}
type car_name struct{
	name string
	ee [4]wheel
}
func main(){
	left_front:=wheel{23.2,32.2}
	left_back:=wheel{22.3,33.2}
	right_front:=wheel{23.5,33.2}
	right_back:=wheel{23.1,23.3}
	fancy_car:=car_name{name:"Nissan",ee:[4]wheel{left_front,left_back,right_back,right_front}}
	fmt.Println("Jeromre",fancy_car.name)
	fmt.Println("left_front:",left_front,"\tright_front",right_back)
	fmt.Println("left_front:",left_back,"\tright_back",right_back)

}
