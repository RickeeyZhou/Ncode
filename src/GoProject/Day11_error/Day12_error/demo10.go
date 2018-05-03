package main

import (
	"fmt"
	"math"
)

func main(){
	i:=-100
	fmt.Println(math.Abs(float64(i)))
	fmt.Println(math.Ceil(5.2))
	fmt.Println(math.Floor(5.22))
	fmt.Println(math.Mod(11,3))
	fmt.Println(math.Modf(5.34))
	fmt.Println(math.Pow(3,2))
	fmt.Println(math.Pow10(4))
	fmt.Println(math.Pi)
}
