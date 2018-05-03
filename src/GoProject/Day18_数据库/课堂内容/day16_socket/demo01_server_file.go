package main

import (
	"net"
	"fmt"
	"os"
)

func main()  {
	//接收图片数据，保存到服务器的本地
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",":9988")
	checkErr(err)
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	checkErr(err)
	fmt.Println("服务器端已经就绪，等待客户端的连接。。")
	conn,err:=listener.Accept()
	checkErr(err)
	fmt.Println("客户端已经连入。。", conn.RemoteAddr())
	//读写操作
	file,err:=os.OpenFile("images\\aa.jpeg",os.O_WRONLY|os.O_CREATE,os.ModePerm)
	//1，从客户端读取数据
	bs := make([] byte,1024)
	for{
		n,err:= conn.Read(bs)
		if n == 0 || err != nil{
			fmt.Println("服务器端已经读取全部数据。。")
			break
		}
		fmt.Println("读到的数据量：",n, err)
		file.Write(bs[:n])
	}
	conn.Close()
	file.Close()



}

func checkErr(err error){
	if err != nil{
		fmt.Println("程序中产生错误了。。",err.Error())
		os.Exit(1) //0表示正常结束，非0表示不正常
	}

}
