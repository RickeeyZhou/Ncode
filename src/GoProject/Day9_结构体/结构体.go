package main

import "fmt"

type dog struct{
	name string
	color string
	age int
}
func main(){
	d1:=dog{"marry","black",3}
	d2:=d1
	d2.name="two_black"
	fmt.Println(d1)
	fmt.Println(d2)
fmt.Println()
	d3:=new(dog)
	d3.name="erer"
	d3.color="yellow"
	d3.age=2

	d4:=d3
	fmt.Printf("%T,%T\n",d3,d4)
	d4.name="big yellow"
	fmt.Println(d3)
	fmt.Println(d4)

	d5:=&d1
	fmt.Printf("%T\n",d5)
	d5.name="黑狗"
	fmt.Println(d1)

}
