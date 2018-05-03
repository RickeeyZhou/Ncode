package main

import "fmt"

type car struct {
	color string
	speed int
}
type bike struct {
	car
	brand string
}
type super_car struct{
	car
	price int
}
func (c car)moving(){

}
func (c car)stop(){

}
func (a bike)footbake(){

}
func (a super_car)trubo(){

}
func main(){
fmt.Println(super_car.stop)
frali:=super_car{car{color:"green",speed:222},222}
fmt.Println(frali)
}
