package main

import (
	"os"
	"fmt"
)

func main(){
	file,_:=os.OpenFile("/home/rickeey/闹着玩.txt",_,_)
	defer file.Close()
	bs:=make([]byte,1)

	file.Seek(4,1)
	count,_:=file.Write(bs)
	fmt.Println(count)
	fmt.Println("i'm done")
}

