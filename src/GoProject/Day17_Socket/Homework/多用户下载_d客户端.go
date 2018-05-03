package main

import (
	"net"
	"os"
	"fmt"
)

func main(){
	tcpAddr,_:=net.ResolveTCPAddr("tcp4","localhost:12345")
	tcpconn,_:=net.DialTCP("tcp4",nil,tcpAddr)
	fmt.Println("客户端连接")
	des,_:=os.OpenFile("/home/rick/周于敬/copy2.jpg",os.O_RDWR|os.O_CREATE,0777)
	bs:=make([]byte,1024)
	for{
		bs=[]byte("")
		count,err:=tcpconn.Read(bs)
		if count==0||err!=nil{
			fmt.Println("接收完毕")
			break
		}
		//tcpconn.Write([]byte(""))
		des.Write(bs[:count])


	}
	tcpconn.Close()
	des.Close()

}
