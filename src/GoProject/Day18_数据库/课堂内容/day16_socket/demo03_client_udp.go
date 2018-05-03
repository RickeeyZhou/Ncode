package main

import (
	"net"
	"fmt"
)

func main()  {
	//基于UDP协议。客户端--->发送方
	udpAttr,_:=net.ResolveUDPAddr("udp4","10.0.154.209:54321")
	udpConn,_:=net.DialUDP("udp",nil,udpAttr)
	udpConn.Write([]byte("面朝大海"))

	udpConn.Close()
	fmt.Println("发送方已经发送数据。。。")
}