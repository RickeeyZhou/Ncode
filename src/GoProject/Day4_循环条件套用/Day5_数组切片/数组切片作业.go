package main

import "fmt"

func main(){
	var a  =[5]int64{23}
	a1:=[...]int{1,23,3,2331,1}
	fmt.Println(a1,a)
	a3:=[3]string{"23","2",","}
	a4:=[...]float64{2.2,232.2,1212.2}

	for index,value:=range a{
		fmt.Print(index,value,"\t")
	}
	fmt.Println()
	for index,value:=range a1{
		fmt.Print(index,value,"\t")
	}

	fmt.Println()
	for index,value:=range a3{
		fmt.Print(index,value,"\t")
	}
	fmt.Println()
	for index,value:=range a4{
		fmt.Print(index,value,"\t")
	}
	fmt.Println()
}