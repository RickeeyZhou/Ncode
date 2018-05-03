package main

import (
	"math"
	"fmt"
)

func main(){
	var B ,S,G float64
	for i:=100;i<=999;i++{
	//i:=153
		B=float64(i/100)
		B=math.Pow(B,3)
		S=float64(i%100/10)
		S=math.Pow(S,3)
		G=float64(i%100%10)
		G=math.Pow(G,3)
		//fmt.Println(B,"\t",S,"\t",G)
		 if G+S+B==float64(i){
		 	fmt.Println(i)
		 }
	}
}
