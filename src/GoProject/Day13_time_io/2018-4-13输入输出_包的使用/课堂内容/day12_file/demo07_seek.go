package main

import (
	"os"
	"fmt"
)

func main()  {
	/*
	seek(offset int, whence int),表示设置光标的位置
		offset int，设置多少个字节
	whence int，从哪开始
		0，距离文件开头
		1，当前的位置
		2，距离文件末尾
	 */
	 //file,_:= os.Open("C:\\Ruby\\pro\\aa.txt") //RDONLY
	 file,_:=os.OpenFile("C:\\Ruby\\pro\\aa.txt",os.O_RDWR,0)
	 //1.打开文件后，光标默认在文件开头。
	 bs := make([] byte, 1)
	 defer file.Close()

	 //2.seek()
	 //设置光标的位置在：距离文件开头，4个字节处。
	 //file.Seek(8,0)
	 count, _:= file.Read(bs)
	fmt.Println(string(bs[:count])) //a

	 file.Seek(4,2)
	//count, _= file.Read(bs)
	//fmt.Println(string(bs[:count])) //
	file.Write([]byte{65,66,67})
	fmt.Println("写完了。。")
}
