package main

import (
	"math"
	"fmt"
)

type point struct{
	x float64
	y float64
	z float64
}
 func distance (a,b point )float64 {
 	var sum  float64
 	sum=math.Sqrt(math.Pow((a.x-b.x),2)+math.Pow((a.y-b.y),2))
 	return sum
 }
func main(){
	var a,b point
fmt.Println("input two point")
fmt.Scanf("%f,%f,%f",&a.x,&a.y,&a.z)
	fmt.Println(a)
fmt.Scanf("%f,%f,%f",&b.x,&b.y,&b.z)
fmt.Println(distance(a,b))
}
