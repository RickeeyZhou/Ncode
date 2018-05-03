package main

import (
	"sort"
	"fmt"
)

func main() {

	//for _, value := range s1 {
	//	fmt.Printf("%q,%T", value, value)
	//}
	//fmt.Println()
	////var map1 map[int32]int=map[int32]int{}

	//	//for j:='A';j<='Z';j++{
	//map2[j]=0
	//}
	//fmt.Println(map2)
	map2 := make(map[int32]int)
	const a = "ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
	s1 := a[:]
	for i := 'A'; i <= 'Z'; i++ {
		for _, value := range s1 {
			if value == i {
				map2[i]++
			}
		}
	}
	//把 键值取出来 排序输出
	var s2 []int
	for key, _ := range map2 {
		s2 = append(s2, int(key)) ///每次 都只有一个元素， s2=[]int{集合} s[]
		//	fmt.Printf("%T:%d\n",key,value)
	}
	sort.Ints(s2)
	for _,key:=range s2{
		fmt.Printf("%c\t%d,\t",key,map2[rune(key)])
	}
	fmt.Println()
	var i byte
	i=65
	fmt.Printf("%c%q",i,i)
	//var s9=[]int {}
	//for i:=0;i<len(s1);i++{
	//	s1[i]
	//}
	//var map1 map[int] int
	//map1=map[int]int{1:2,}
	//var map2:= map[int]int{2:2,3:1}
	//map2[key]=2

	//	fmt.Printf("E:%d",E)
arr5:="12221G"
a11:=arr5[:]
a112:=[]byte(arr5)
fmt.Println(a11,arr5[2])
fmt.Printf("%T",arr5[5])
sss:=[]int32(arr5)
fmt.Printf("\n%T",sss)
}
