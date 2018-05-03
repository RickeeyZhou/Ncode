package main

import (
	"net"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main()  {
	tcpAddr,_:=net.ResolveTCPAddr("tcp4","10.0.154.209:6677")
	tcpConn,_:=net.DialTCP("tcp",nil,tcpAddr)
	fmt.Println("客户端已经连入。。")
	bs := make([] byte,1024)
	file,_:=os.OpenFile("images\\aa"+strconv.FormatInt(time.Now().UnixNano(),10)+".jpeg",os.O_WRONLY|os.O_CREATE,os.ModePerm)
	for{
		count,err:=tcpConn.Read(bs)
		if count ==0||err != nil{
			fmt.Println("客户端已经下载完毕。。")
			break
		}
		file.Write(bs[:count])
	}
	file.Close()
	tcpConn.Close()
}
