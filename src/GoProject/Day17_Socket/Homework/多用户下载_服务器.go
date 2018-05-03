package main

import (
	"net"
	"os"
	"io"
	"fmt"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", ":12345")
	lisentener, _ := net.ListenTCP("tcp4", tcpAddr)
	for {
		fmt.Println("wait conn\n")
		conn, _ := lisentener.Accept()
		fmt.Println("连接成功")
		go copy(conn)
	}
}
func copy(conn net.Conn) {
	src, _:= os.Open("/home/rick/周于敬/me.jpeg")
	bs := make([]byte, 1024)
	for {
		count, err := src.Read(bs)
		if err == io.EOF {
			//conn.Write(bs[:count])
			//conn.Read(bs)
			break
		}
		conn.Write(bs[:count])
	}
	src.Close()
	conn.Close()

}
