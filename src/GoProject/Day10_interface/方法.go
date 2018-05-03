package main

import "fmt"

func main(){
	FF:=car {
		name:"fff",
		speed:222,
	}
	FF.Pspeed(FF.name)
}
type car struct {
	name string
	speed int

}
func (c car )Pspeed(a string){
	fmt.Println(c.speed,"km/h")
}