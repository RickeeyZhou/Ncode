package main

import "fmt"

func main(){
	fmt.Println( dv(2 ,0))
}
func dv(a int,b int )(int,int){
	defer func (){
		recover()
	}()
	defer recover()
	if b==0{
		panic("分母不能为0")
}
	x:=a/b
	y:=a%b
	return x,y
}