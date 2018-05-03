package main

import (
	"os"
	"fmt"
)

func main() {
	file, _ := os.OpenFile("/home/rickeey/测试/ToDoList.sh", os.O_RDWR, 0)
	defer file.Close()

	bs := make([]byte, 1)

	count, _ := file.Read(bs)
	fmt.Println(string(bs[:count]))
	file.Seek(4, 1)
	file.Write([]byte{22,31,3})
	fmt.Print("i,m done")
}
