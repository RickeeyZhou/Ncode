package main

import (
	"fmt"
	"sort"
)

func main(){
var  sli []int
fmt.Println("plz input a group num")
var i,  a int
for ;;{
	fmt.Scan(&a)
	sli=append(sli,a )
	i++
	if i>5{
		break
	}
}
fmt.Println(sli)
fuck_u_off(sli )
}
func fuck_u_off(sli1 []int){
	sort.Ints(sli1)
	fmt.Println(sli1)
}
