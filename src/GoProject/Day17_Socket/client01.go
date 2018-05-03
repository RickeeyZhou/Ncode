package main

import (
	"net"
	"fmt"
)

func main(){
	service:=":052022"
	SerName,err:=net.ResolveTCPAddr("tcp4",service)
	fmt.Println(SerName,err)

	tcpConn,err:=net.DialTCP("tcp" ,nil,SerName)
	fmt.Println(tcpConn,err)
	fmt.Printf("%T\n",tcpConn)
	fmt.Println("客户端连接成功")
	fmt.Println("服务器地址：",tcpConn.RemoteAddr())

	//
	count,err:=tcpConn.Write([]byte("fuck you"))
	fmt.Println(count,err)
	fmt.Println("传输完成")

	bs:=make([]byte ,1024)
	n,_:=tcpConn.Read(bs)
	fmt.Println(string(bs[:n]))
	tcpConn.Close()

}
