package main

import (
	"os"
	"fmt"
)

func main(){
	file,err:=os.OpenFile("/home/rickeey/闹着玩.txt",os.O_CREATE|os.O_RDWR,0777)
	fmt.Println(file)  //
	ba:=make([]byte,222)
	count,err:=file.ReadAt(ba,3)
	fmt.Println(err)
	fmt.Println(count)
	fmt.Printf("\n%v,%T",ba[0:count],ba[0:count])
	fmt.Println(string(ba[0:count]))

}
