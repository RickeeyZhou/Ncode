package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	//step1.定义服务端的端口：
	ip := ":9527"
	// step2：获取服务器端的地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip)
	checkErr(err)
	fmt.Printf("%T\n", tcpAddr) //*net.TCPAddr
	fmt.Println(tcpAddr)        // :9527
	// step3：服务端需要监听该端口
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	fmt.Printf("%T\n", listener) //*net.TCPListener
	fmt.Println(listener)        //&{0xc420090000}
	// step4：等待客户端的链接
	fmt.Println("等待客户端的链接。。")

	//clients := make([] *net.TCPConn, 0)

	map1 := make(map[string]net.Conn)
	done:=make(chan bool)
	for i := 0; i < 2; i++ {
		conn, _ := listener.Accept()
		// 获取远程客户端的信息
		remoAddr := conn.RemoteAddr()
		fmt.Println("已有客户端连入：", i,remoAddr) //192.168.1.128:50577

		map1[remoAddr.String()] = conn
		fmt.Println(map1)
		go sendData3(conn, map1,done)
	}
	<-done
}

func sendData3(conn net.Conn, map1 map[string]net.Conn, done chan bool) {
	bs := make([] byte, 512)
	sendName := conn.RemoteAddr().String()
	for {
		count, _ := conn.Read(bs)
		content := string(bs[:count])
		fmt.Println(sendName, "发来消息：", content)
		if content == "over" {
			//conn.Close()
			fmt.Println(sendName, "已经退出。。")
			delete(map1, sendName)
			fmt.Println(map1)
			break
		}

		for key, v := range map1 {
			if key == sendName {
				continue
			}
			msg := sendName + "说：" + content

			v.Write([] byte(msg))
		}

	}
	if len(map1) ==0{
		done<-true
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println("发生了错误。。")
		os.Exit(1) //程序退出，非0表示不正常退出
	}
}
