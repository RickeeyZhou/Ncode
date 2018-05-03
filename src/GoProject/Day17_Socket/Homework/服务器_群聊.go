package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	ip := ":9900"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip)
	checkErr1(err)
	fmt.Printf("%T\n", tcpAddr)
	fmt.Println(tcpAddr)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	fmt.Println(listener)
	fmt.Printf("%T\n", listener)
	fmt.Println("客户端等待连接")
	// client:=

	map1 := make(map[string]net.Conn)
	for {
		con, _ := listener.Accept()
		name := con.RemoteAddr().String()

		map1[name] = con
		fmt.Println(name, "加入聊天")

		go chat(con, map1)

	}

}
func chat(conn net.Conn, map1 map[string]net.Conn) {
	bs := make([]byte, 1024)

	for {
		msg := ""
		count, timeout:= conn.Read(bs)
		if timeout!=nil{
			break
		}
		fmt.Println(conn.RemoteAddr().String(), "say", string(bs[:count]))

		//逐个发送
		for name, add := range map1 {
			if name == conn.RemoteAddr().String() {
				continue
			}
			//判断是否退出聊天室
			if string(bs[:count]) == "over" || string(bs[:count]) == "" {
				// 切记 切片这里 要完整的写 bs[:count]
				fmt.Println(conn.RemoteAddr().String(), "离开聊天室")
				msg = conn.RemoteAddr().String() + "离开聊天室"
				delete(map1, conn.RemoteAddr().String())
				//add.Write([]byte(msg))
			} else {
				msg = conn.RemoteAddr().String() + "say:" + string(bs[:count])
			}
			// 发送
			add.Write([]byte(msg))
		}
	}
}

func checkErr1(err error) {
	if err != nil {
		fmt.Println("发生了错误。。")
		os.Exit(1) //程序退出，非0表示不正常退出
	}
}
