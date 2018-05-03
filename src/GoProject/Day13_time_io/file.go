package main

import (
	"os"
	"fmt"
)

func main(){
	fileInfo,err:=os.Stat("text.go")
	fmt.Println(err)
	fmt.Println("%T\n",fileInfo)
	if err!=nil{
		fmt.Println("doc operation have mistake ",err)
		return
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Mode())
	//err1:=os.Mkdir("//home//rickeey//aa",0666)
	//fmt.Println(err1)
	//err2:=os.Remove("/home/rickeey/aa")
	//fmt.Println(err2)
	mk:=os.MkdirAll("/home/rickeey/zyj2j",111)
	fmt.Println(mk)
	fmt.Println(os.Remove("/home/rickeey/zyj2j"))
	fmt.Println(os.Remove("/home/rickeey/zyjj"))
}