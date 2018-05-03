package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main()  {
	//读取本地文件，传递给服务器
	tcpAddr,_:=net.ResolveTCPAddr("tcp4","10.0.154.209:9988")
	tcpConn,_:=net.DialTCP("tcp",nil,tcpAddr)
	fmt.Println("客户端已经成功连接。。")
	//读写
	//1.打开客户端文件，准备读取
	file,_:=os.Open("C:\\Ruby\\pro\\aa.jpeg")
	bs :=make([] byte,1024)

	for{
		count,err := file.Read(bs)
		if err == io.EOF{
			fmt.Println("客户端数据已经全部读取完毕。。")
			break
		}
		tcpConn.Write(bs[:count])
	}
	tcpConn.Close()
	file.Close()

}
