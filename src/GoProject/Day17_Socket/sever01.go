package main

import (
	"fmt"
	"net"
)

func main(){
	service:=":052022"
	serName,err:=net.ResolveTCPAddr("tcp4",service)
	fmt.Println(serName,err)// 我的服务器名字
	fmt.Printf("%T\n",serName)

	//监听端口
	listener,err:=net.ListenTCP("tcp",serName)
	fmt.Println("监听器是：",listener)
	fmt.Printf("%T\n",listener)

	//等待接入，输入监听器信息
	fmt.Println("server is ready, wait client to come ")
	GetConn,err:=listener.Accept()
	fmt.Println(GetConn)//存储客户端信息的 地址
	fmt.Printf("%T\n",GetConn)
	fmt.Println("connect succuess \n\n\n\n")

	//交换数据
	bs:=make([]byte,1024)
	fmt.Println("已经创建服务器接收容器")
	count,err:=GetConn.Read(bs) //从监听器信息中心读取信息
	fmt.Println("接收数据量：",count)
	fmt.Println("客户端信息整理：",string(bs))

	GetConn.Write([]byte("接收完成"))

	GetConn.Close()





}