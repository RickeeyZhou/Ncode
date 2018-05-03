package main

import "fmt"

func main(){
	var num int
	fmt.Scanln(&num)
	fmt.Println(jie(num))

}
func jie(num int )int {
	if num==1{
		return 1
	}
	return num*jie(num-1)
}