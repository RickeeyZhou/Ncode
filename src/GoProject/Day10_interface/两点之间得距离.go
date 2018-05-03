package main

import (
	"math"
	"fmt"
)

func main(){
	var a point =point{2,3,1}
	b:=point{11,3,-2}
	dis:=distance(a,b)
	fmt.Printf("%.3f",dis)
}
type point struct {
	x float64
	y float64
	z float64
}
func distance(a,b  point )float64{
	var dis float64
	dis=math.Sqrt(  math.Pow((a.x-b.x),2) +math.Pow((a.y-b.y),2)+math.Pow((a.z-b.z),2))
	return dis
}
