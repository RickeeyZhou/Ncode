package main

import (
	"net"
	"fmt"
)

func main() {
	//1.提供本机的地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "10.0.154.238:9527")
	//2.提供监听对象
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	//3.监听端口，等待客户端链接
	fmt.Println("服务器端已经就绪，等待客户端链接。。。")
	conn, _ := listener.Accept()
	defer conn.Close()
	//4.读写数据
	//C：服务器读取客户端
	for {
		bs := make([] byte, 512)
		n, _ := conn.Read(bs) //阻塞的
		content:=string(bs[:n])
		if content=="over" ||n == 0{
			fmt.Println("客户端已经跑了。。")
			break
		}
		fmt.Println("客户端说：", content, n)
		//D：服务器读键盘
		line := ""
		fmt.Scanln(&line) //阻塞式
		//E：写给客户端
		conn.Write([] byte(line))
	}
}

/*
练习1：客户端和服务端：
	多句聊天，直到客户端说over。
客户端-->服务器：
	你干啥呢
服务器-->客户端：
	我等你呢
客户端-->服务器
	啥事

练习2：ECHO程序
	客户端读取本地键盘，发送给服务器："啊啊啊啊啊"
	服务器，接收到数据后，添加ECHO，发送给客户端：
		"ECHO:啊啊啊啊啊""
	直到客户端发送over为止。
 */
