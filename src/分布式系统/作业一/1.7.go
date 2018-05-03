package main

import "fmt"

func main(){
	s:=make([]int,5)
	s=append(s,1,2,3)
	fmt.Println(s)

}
//
//{1,2,3,0,0}
//{0,0,0,0,0,1,2,3}
//创建切片,一旦创建长度,就意味着有了实体,就会附上默认值
//使用make创建切片时,会内部会分配数组,返回对应的切片值