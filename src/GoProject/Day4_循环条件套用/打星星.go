package main

import (
	"fmt"
)
 func main(){
 	for i:=0;i<5;i++{
 		for j:=1;j<=9;j++{
 			if 5-i<j&&j<5+i{
 				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

 }

