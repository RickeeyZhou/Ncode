package 接口

import "fmt"


type Car struct {
	Color string
	Speed int
}
func (a Car )moving (){
	fmt.Println(" Moving as a car..")
}
func (a Car )stop(){
	fmt.Println("Stop as a car ..")
}
type Bike struct{
	Car
	Price int
}
type Super_car struct{
	Car
	Car_brand string
}
func (a Bike )visiting(){
	fmt.Println("visiting as a bike ")
}
func (a Super_car)visiting(){
	fmt.Println("visiting as a super_car")
}
