package main

import (
	"os"
	"fmt"
)

func main(){
	file,_:=os.Open("/home/rickeey/闹着玩/深度截图_选择区域_20180413074933.png")
	//defer file.Close()

	bs:=[]byte{22,3,23,2,2,3,2}
	file.Write(bs[:3])
	ba:="fuck you"
	file.WriteString(ba)
	err:=os.Remove("/home/rickeey/截图/深度截图_选择区域_20180413074933.png")
	fmt.Println(err)


}
