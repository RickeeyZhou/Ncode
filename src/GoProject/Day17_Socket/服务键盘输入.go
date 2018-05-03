package main
////服务器
import (
	"net"
	"fmt"
)

func main()  {
	//1.提供本机的地址
	tcpAddr,_:=net.ResolveTCPAddr("tcp4","10.0.154.200:9522")
	//2.提供监听对象
	listener,_:=net.ListenTCP("tcp",tcpAddr)
	//3.监听端口，等待客户端链接
	fmt.Println("服务器端已经就绪，等待客户端链接。。。")
	conn,_:=listener.Accept()
	//defer conn.Close()
	//4.读写数据
	//C：服务器读取客户端
	bs :=make([] byte,512)
	for {
		n, _ := conn.Read(bs) //阻塞的

		fmt.Println("客户端说：", string(bs[:n]))
		if string(bs[:n])=="over"{
			break
		}
		//D：服务器读键盘
		line := ""
		fmt.Scanln(&line) //阻塞式
		//E：写给客户端
		conn.Write([] byte(line))

		if string(bs[:n]) == "over" {
			conn.Write([] byte("over"))
			return
		}

	}
}