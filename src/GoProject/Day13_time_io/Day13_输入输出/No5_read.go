package main

import (
	"os"
	"fmt"
	"io"
)

func main(){
	file,err:=os.Open("/home/rickeey/测试/ToDoList.sh")
	if err!=nil{
		fmt.Print("something wrong!!\n")
		return
	}
	defer file.Close()
	bs:=make([]byte,4)
	count:=0
	for{
		count,err=file.Read(bs)
		if err==io.EOF{
			fmt.Println()
			fmt.Println("get data over ")
			break
		}
		fmt.Print(string(bs[:count]))
	}
}