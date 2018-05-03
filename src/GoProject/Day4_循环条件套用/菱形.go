package main
/*
    *
   ***
  *****
 *******
*********

 *******
  *****
   ***
    *
第一层循环 5次
第二层循环 9次
 */
import "fmt"

func main() {
	for i:=1;i<=5;i++{
		for j:=1;j<=5-i;j++{
			fmt.Print(" ")
		}
		for j:=1;j<=2*i-1;j++{
			fmt.Print("*")
		}

		fmt.Println()
	}


}

