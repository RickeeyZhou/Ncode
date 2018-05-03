package main

import (
	"net"
	"os"
	"io"
)

func main(){
serverName:="10.0.154.200:9522"
tcpAddr,_:=net.ResolveTCPAddr("tcp4",serverName)

tcpConn,_:=net.DialTCP("tcp" ,nil,tcpAddr)
defer tcpConn.Close()
pic,_:=os.OpenFile("/home/rick/周于敬/进大学前.jpg",os.O_RDONLY,0777)
for {

	bs := make([]byte, 512)
	count,err:= pic.Read(bs)
	if err==io.EOF{
		tcpConn.Write(bs[:count])
		break
	}
	tcpConn.Write(bs[:count]) //产生阻塞
}
tcpConn.Close()

}
