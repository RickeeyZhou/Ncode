package main
/////客户端
import (
	"net"
	"fmt"
)

func main()  {
	//1.提供服务器地址
	tcpAddr,_:=net.ResolveTCPAddr("tcp4","10.0.154.200:2018" )
	//2.连接
	tcpConn,_:=net.DialTCP("tcp",nil,tcpAddr)
	defer tcpConn.Close()
	//3.读写
	//A.读取键盘
	line:=""
	flag:fmt.Scanln(&line)//键盘输入，阻塞的
	//B.写给服务器
	tcpConn.Write([] byte(line))
	//F:读取服务器
	bs:=make([] byte,512)
	n,_:=tcpConn.Read(bs)//阻塞式
	fmt.Println("服务器说：",string(bs[:n]))
	if string(bs[:n])=="over"{

		return
	}
	goto flag
	//关闭


}
