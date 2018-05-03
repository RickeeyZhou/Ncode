package main

import (
	"net"
	"fmt"
	"os"
)

func main(){
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",":9900")
	checkErr(err)

	tcpConn,err:=net.DialTCP("tcp4",nil,tcpAddr)
	fmt.Printf("%T\n",tcpConn)
	done:=make(chan bool)
	go sendDATA(tcpConn,done)
	go getDATA(tcpConn )
	<-done
}
func checkErr(err error){
	if err!=nil{
		fmt.Println("地址解析错误",err)
		os.Exit(1)
	}

}
func sendDATA(conn *net.TCPConn,done chan bool){
	var string1 string
	for {
		fmt.Scanln(&string1)
		conn.Write([]byte(string1))
		fmt.Println("传送成功")
		if string1==""||string1=="over"{
			done<-true
			os.Exit(0)
		}
	}
}
func getDATA(conn *net.TCPConn){
	bs:=make([]byte,1024)
	for {
		count,_:=conn.Read(bs)
		if count==0{
			fmt.Println("exit" )
			break
		}
		fmt.Println(string(bs[:count]))

	}
}
