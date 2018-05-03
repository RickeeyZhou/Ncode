package main

import "fmt"

func main(){
	// 最后一天剩下x个桃子
	x:=1 //10
 	//x=tutaozi(x ) //9
 	//x=tutaozi(x)//8
 	for i:=10;i>=1;i--{
 		x=tutaozi(x)
	}
	fmt.Println(x)

}
func tutaozi (x int )int{
	  return (x+1)*2
}