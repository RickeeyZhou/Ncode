package main

import (
	"math"
	"fmt"
)

func main(){
tri:=triangle{3,3,3}
cir:=circle{2}
ort:=orthogon{3,1.4}
fmt.Println("三角周长：",tri.area(),"三角面积：",tri.area())
fmt.Println("圆：",cir.area(),"三角面积：",cir.area())
fmt.Println("矩形周长：",ort.area(),"矩形面积：",ort.area())
testprimeter(ort)
}
func testprimeter(a apperence){
	fmt.Println(a.area(),a.perimeter())
}
type apperence interface {
	  perimeter()float64
	  area()float64
}
type triangle struct{
	a float64
	b float64
	c float64
}
type orthogon struct{
a float64
b float64
}
type circle struct {
r float64
}
func (a triangle) perimeter()float64{
sum:=a.b+a.a+a.c
return sum
}
func (a orthogon) perimeter()float64{
	sum :=2*a.a+2*a.b
	return sum
}
func (a circle )perimeter ()float64 {
	sum:=2*3.14*a.r
	return sum

}
func (a triangle )area()float64{
	per:=a.perimeter()/2
	sum:=math.Sqrt(per*(per-a.a)*(per-a.b)*(per-a.c))
	return sum
}
func (a orthogon)area()float64{
	sum:=a.b*a.a
	return sum

}
func (a circle)area()float64{
	sum:=3.14*math.Pow(a.r,2)
	return sum
}