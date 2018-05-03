package main

import "os"

func main(){
	file,_:=os.OpenFile("/home/rickeey/测试/ToDoList.txt",os.O_RDWR,0)
	defer file.Close()

	bs:="fuck you you "
	file.WriteString(bs)




}
