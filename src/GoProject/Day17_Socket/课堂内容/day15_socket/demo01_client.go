package main

import (
	"net"
	"fmt"
)

func main()  {
	//TCP的客户端程序
	//step1:提供要连接的服务端的地址
	service:="10.0.154.238:54321" //是服务器的地址，不是自己的，自己的端口由系统自动分配
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",service)
	fmt.Println(tcpAddr,err) //10.0.154.238:54321 nil
	// step2：申请连接服务器
	tcpConn, err:=net.DialTCP("tcp",nil,tcpAddr)
	fmt.Println(tcpConn,err) //&{{0xc042088000}} <nil>
	fmt.Printf("%T\n", tcpConn) //*net.TCPConn
	fmt.Println("客户端已经连接成功。。。")
	fmt.Println("服务器的地址：",tcpConn.RemoteAddr())
	//step3：数据交互
	n,err:=tcpConn.Write([] byte("hello"))//写出的字节的数量
	fmt.Println(n,err)
	fmt.Println("数据已经写出。。")

	bs := make([] byte,512)
	n,err=tcpConn.Read(bs)
	fmt.Println(string(bs[:n]))
	//step4.关闭资源，断开连接
	tcpConn.Close()//
}
/*
客户端是你自己
客户端读取键盘输入，传递给服务端一句话
服务端是你同桌
服务端读取键盘输入，传递给客户端一句话
 */



