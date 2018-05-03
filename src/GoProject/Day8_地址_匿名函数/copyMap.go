package main

import "fmt"

func main(){
//map copy
map1:=map [int]int{1:2,2:3,4:2,5:22}
fmt.Println(map1 )
map2:=map1
fmt.Println(map2)
map1[1]=3
fmt.Println(map1)
fmt.Println(map2)
/// shallow copy
var sli1 []int=[]int {}
sli2:=make([]int,0)
for key,value :=range map1{
	sli2=append(sli2,value)
	sli1=append(sli1,key)
}
map3:=make (map [int ]int )
for i:=0;i<len(sli1);i++{
	map3[sli1[i]]=sli2[i]
}
fmt.Println(map1)
fmt.Println(map3)
for key,value:=range map1{
	map1[key]=value+1
}
fmt.Println(map1)
fmt.Println(map3)
}
