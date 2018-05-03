package main

import (
	"net"
	"fmt"
)

func main() {
	//1.提供服务器地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "10.0.154.238:9527")
	//2.连接
	tcpConn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer tcpConn.Close()
	//3.读写
	for {
		//A.读取键盘
		line := ""
		fmt.Scanln(&line) //键盘输入，阻塞的
		//fmt.Println("--",line,"--")
		//B.写给服务器
		n,_:=tcpConn.Write([] byte(line))
		fmt.Println("客户端写出数据：",n)
		if line == "over"{
			fmt.Println("客户端即将结束。。")
			break
		}
		//F:读取服务器
		bs := make([] byte, 512)
		n,_= tcpConn.Read(bs) //阻塞式
		fmt.Println("服务器说：", string(bs[:n]))
	}
	//关闭

}
