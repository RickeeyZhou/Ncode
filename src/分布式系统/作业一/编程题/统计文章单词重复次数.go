package main

import (
	"strings"
	"fmt"
)

func main(){
	//以空格为分割符
	map1:=make(map[string]int)
	s:=" adadada sddd sdadq sdadadad"
	s1:=strings.Split(s," ")
	fmt.Printf("%T\n",s1)
	for _,v :=range s1{
		map1[v]=strings.Count(s,v)
	}
	fmt.Println(map1)
}
