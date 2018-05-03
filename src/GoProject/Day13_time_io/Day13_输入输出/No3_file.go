package main

import (
	"os"
	"fmt"
)

func main(){
	fileInfo,err:=os.Stat("/home/rickeey/测试/ToDoList.sh")
	fmt.Println(err)

	fmt.Println(fileInfo)
	fmt.Println("%T\n",fileInfo)
	if err!=nil{
		fmt.Println()
		fmt.Println("doc file opration have mistake",err)
	return
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Mode())



}
