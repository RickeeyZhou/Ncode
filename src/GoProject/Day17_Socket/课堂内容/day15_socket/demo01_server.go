package main

import (
	"net"
	"fmt"
)

func main()  {
	//基于TCP的服务器
	//step1：确定本机的地址：ip:port--->TCPAddr烈性
	service := ":54321" // string
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",service)
	fmt.Println(tcpAddr,err)
	fmt.Printf("%T\n",tcpAddr) //*net.TCPAddr

	//step2：监听该端口
	listener,err:=net.ListenTCP("tcp", tcpAddr)
	fmt.Println(listener,err)
	fmt.Printf("%T\n",listener) //*net.TCPListener
	//step3：接收客户端的连接请求
	fmt.Println("服务器程序已经就绪，等待客户端的链接。。。")
	conn,err:=listener.Accept()//阻塞式
	fmt.Println(conn, err) //&{{0xc042076000}} <nil>
	fmt.Printf("%T\n", conn) //*net.TCPConn
	fmt.Println("已有客户端连入。。",conn.RemoteAddr())

	//step4：数据交互
	bs := make([] byte,512)
	n,err:=conn.Read(bs)
	fmt.Println(n,err)
	fmt.Println("客户端说：", string(bs[:n]))
	conn.Write([] byte("我是服务器"))

	//step5：关闭资源
	conn.Close()
}
