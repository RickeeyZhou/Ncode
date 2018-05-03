package main

import (
	"net"
	"fmt"
)

func main()  {
	//1UDP协议，接收方
	udpAttr,_:=net.ResolveUDPAddr("udp4",":54321")
	conn,_:=net.ListenUDP("udp",udpAttr)
	bs := make([] byte, 512)
	//n,_:=conn.Read(bs)
	n, rudpAttr,_:=conn.ReadFromUDP(bs)
	fmt.Println("接收到的数据是：",string(bs[:n]),"发送方是：",rudpAttr)
	conn.Close()
}
