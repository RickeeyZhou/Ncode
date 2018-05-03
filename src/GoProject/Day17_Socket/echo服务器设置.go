package main

import (
	"net"
	"fmt"
)

func main(){
	server:="10.0.154.200:2018"
	tcpAddr,_:=net.ResolveTCPAddr("tcp4",server)

	listener,_:=net.ListenTCP("tcp4",tcpAddr)

	fmt.Println("服务器架设完毕，等待客户端开机")


	for {
		conn, _ := listener.Accept()
		 go change(conn)
	}
}

func change(conn net.Conn ) {
	for {
		bs := make([]byte, 512)
		n, _ := conn.Read(bs)
		fmt.Println("成功接收数据，正在Echo", bs)
		//cotent := []byte("")
		cotent := []byte("echo:" + string(bs[:n])) ////////****
		conn.Write(cotent)
		fmt.Println(cotent)
		fmt.Println(conn.RemoteAddr())
		cotent = []byte("")
	}
}