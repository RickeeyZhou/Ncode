package main

import (
	"net"
	"fmt"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", ":8899")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("服务器已经就绪，等待客户端链接。。")
	for {
		conn, _ := listener.Accept()
		fmt.Println("已有客户端连入。。", conn.RemoteAddr())
		bs:=make([] byte,512)
		n,_:=conn.Read(bs)
		fmt.Println("接收到的数据是：", string(bs[:n]))
		conn.Close()
	}
}
