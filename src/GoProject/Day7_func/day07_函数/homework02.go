package main

import (
	"fmt"
	"strings"
)

func main()  {
	/*
	aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com
	 */
	 str:= "aa:zhangsan@163.com!bb:lisi@sina.com!cc:wangwu@126.com"

	 arr1:= strings.Split(str,"!")
	 map1:= make(map[string]string)
	 for i:=0;i<len(arr1);i++{
	 	fmt.Println(arr1[i])
	 	//arr2:= strings.Split(arr1[i],":")
	 	//map1[arr2[0]] = arr2[1]
	 	index:=strings.Index(arr1[i],":")
		map1[arr1[i][:index]] = arr1[i][index+1:]
	 }
	 fmt.Println(map1)

	pathName:="http://192.168.10.1:8080/Day33_Servlet/aa.jpeg"
	index2:=strings.LastIndex(pathName,"/")
	fileName:=pathName[index2+1:]
	fmt.Println(fileName)
}
