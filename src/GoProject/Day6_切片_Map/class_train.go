package main

import "sort"
import (
	"fmt"
)

func main() {
	var map1 map[int]string
	map1 = map[int]string{
		1: "goer",
		2: "hr",
		3: "inviter",
		4: "Cto"}
	var arr []int
	for key := range map1 {
		arr = append(arr, key)
	}
	sort.Ints(arr)
	fmt.Println(arr)

	//for key,value:=range map1{
	//
	//}
	fmt.Print(map1[arr[2]],"\n")
	fmt.Print(map1)
	//

}