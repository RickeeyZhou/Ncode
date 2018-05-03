package main

import (
	"fmt"
	"math"
)

func main(){
	// 求出 每一项,一共有A 项
	var a int
	var sum int
	var value int
	fmt.Scanln(&a )
	for  i:=1;i<=a ;i++{

		// 一共有a 位
		for j:=1;j<=i;j++{
 			if j==1 {
				value=a
			}else{
				value = value+a *int( math.Pow10(j-1))
			}
			//fmt.Println(value)
		}

		sum=sum+value
	}

	fmt.Println(sum)

}
