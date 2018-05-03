package main

import (
	"net"
	"os"
	"fmt"
)

func main(){
	//server:="10.0.154.200:2995"
	tcpAddr,_:=net.ResolveTCPAddr("tcp4","10.0.154.200:9522")

	listener,_:=net.ListenTCP("tcp",tcpAddr)
	serConn,_:=listener.Accept()
	desname,_:=os.OpenFile("/home/rick/周于敬/copy.jpeg",os.O_RDWR|os.O_CREATE,0777)
	defer desname.Close()
	bs:=make([]byte,512)

for {
	fmt.Println("开始发送")
	count, _ := serConn.Read(bs)
	//fmt.Println(serConn.RemoteAddr())
	desname.Write(bs[:count])

	serConn.Close()
}


}
