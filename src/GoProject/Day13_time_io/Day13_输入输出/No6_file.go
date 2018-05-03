package main

import (
	"os"
	"fmt"
)

func main(){
	file,_:=os.Open("/home/rickeey/测试/ToDoList.sh")
	defer file.Close()

	bs:=make([]byte,4)

	count,err:=file.ReadAt(bs,2)
	fmt.Println(err)
	fmt.Println(count)
	fmt.Println(string(bs[:count]))

	count,err=file.ReadAt(bs,7)
	fmt.Println(err)
	fmt.Println(count)
	fmt.Println(string(bs[:count]))





}
