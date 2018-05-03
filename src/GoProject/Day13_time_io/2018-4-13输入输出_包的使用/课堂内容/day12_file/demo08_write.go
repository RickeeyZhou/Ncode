package main

import (
	"os"
	"fmt"
)

func main()  {
	/*
	Writer(),向外写数据
	 */
	 file,_:=os.OpenFile("C:\\Ruby\\pro\\ao.txt",os.O_WRONLY|os.O_CREATE,0777)
	 defer file.Close()
	 //直接写出字节数组、



	 bs := [] byte{97,98,99,100,101,102} // a-f
	 file.Write(bs[0:4]) //0123
	file.WriteString("面朝大海")
	//abcd面朝大海
	/*
	16个字节
	1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
	a b c d   面    朝      大      海
	abcd面朝大ABCDE
	 */


	file.WriteAt([]byte{'A','B','C','D','E'},13)
	 fmt.Println("写完啦啊。。")
/*
练习1：readAt(),seek(),熟悉用法
练习2：write(byte[]), writestring(s),writeAt()
练习3：复制文件
	make([]byte, 1024)
	C:\\Ruby\\pro\\aa.jpeg,
	D:\\copy\\aa.jpeg
 */




}
