package main

import (
	"net"
	"fmt"
	"os"
)

func main()  {
	//step1.确定要访问的服务端的ip和端口
	service := "127.0.0.1:9527"
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",service)
	fmt.Printf("%T\n",tcpAddr) //*net.TCPAddr
	fmt.Println(tcpAddr) //192.168.1.128:9527
	checkErr(err)
	// step2：客户端申请链接服务器
	tcpConn, err:=net.DialTCP("tcp4",nil,tcpAddr)
	fmt.Printf("%T\n",tcpConn) //*net.TCPConn
	fmt.Println(tcpConn) //*net.TCPConn


	// step3：向服务器发送消息

	go sendData(tcpConn)


	line:=""
	for{
		fmt.Scanln(&line)
		tcpConn.Write([] byte(line))
		fmt.Println("客户端数据写完。。")
		if line==""||line=="over"{
			fmt.Println("客户端退出程序。。")
			break
		}

	}
	tcpConn.Close()


}
func sendData(tcpConn *net.TCPConn){
	bs:=make([] byte, 512)
	for{
		count,_:=tcpConn.Read(bs)
		if count ==0{
			fmt.Println("退出。。")
			break
		}
		fmt.Println(string(bs[:count]) ,count)
	}
}





func checkErr(err error){
	if err != nil{
		fmt.Println("错误了。。", err)
		os.Exit(0)
	}
}
