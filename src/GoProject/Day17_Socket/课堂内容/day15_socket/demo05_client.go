package main

import (
	"net"
	"fmt"
)

func main()  {
	tcpAttr,_:=net.ResolveTCPAddr("tcp4","10.0.154.238:8899")
	tcpConn,_:=net.DialTCP("tcp",nil,tcpAttr)
	fmt.Println("客户端已经连接成功。。", tcpConn.RemoteAddr())
	line:=""
	fmt.Scanln(&line)
	tcpConn.Write([] byte(line))
	tcpConn.Close()
}
