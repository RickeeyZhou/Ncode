package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main()  {
	tcpAddr, _ :=net.ResolveTCPAddr("tcp4",":6677")
	listener,_:=net.ListenTCP("tcp",tcpAddr)
	fmt.Println("服务器端已经建立。。")
	for{
		//1.接收客户端的链接
		conn,_:=listener.Accept()
		//2.启动子goroutine，给客户端传递数据
		go func() {
		//假装c盘pro目录是服务器本地盘符，aa.jpeg
			file,_:=os.Open("C:\\Ruby\\pro\\aa.jpeg")
			bs:=make([] byte,1024)
			for{
				count,err:=file.Read(bs)
				if err == io.EOF{
					fmt.Println("服务器端数据全部传递完毕。。")
					break
				}
				conn.Write(bs[:count])
			}
			conn.Close()
			file.Close()

		}()
	}
}
