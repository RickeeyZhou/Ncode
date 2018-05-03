package main

import (
	"os"
	"io"
	"fmt"
)

()
import("os"
		"fmt"
		"strconv"
		"io"
)
func main(){
	/*
	断点续传
	采用临时文件 储存断点的信息，我之前没有想到临时文件，以为变量可以储存，
	其实只要不写，数据在内存中都会马上消失，
	问题的关键在于创建文件，传输文件，并且记录每一个传输时刻
	我的原来的想法，根本就是在用变量total记录，一旦停止，变量就没了，
	我对”停止“这件事根本就没有理解，也不知道停止会发生什么，程序停止，相当于没有这个程序
	程序运行，在内存中跑起来，一旦停止内存清空，相当于没有发生，
	在这题中，输入输出IO 在文件中的操作经保留下来
	 */
	 srcName:="/home/rickeey/测试/ToDoList.sh"
	 destName:="/home/rickeey/测试/ToDoList.sh"
	 file1,_:=os.Open(srcName)
	 file2,_:=os.OpenFile(destName,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	 tempName:=destName+"temp.txt"
	 file3,_:=os.OpenFile(tempName,os.O_CREATE|os.O_RDWR,os.ModePerm)
	 defer file1.Close()
	 defer file2.Close()
	 //
	 var total int

	 //常规拷贝
	 file1.Seek(int64(total),0)
	 file2.Seek(int64(total),0)
	 bs:=make([]byte,1024)
	 for {
		 count2, err := file1.Read(bs)
		 if err== io.EOF{
		 	fmt.Println("the end..",count2)
		 	file2.Write(bs[:count2])
		 	break
		 }
		 file2.Write(bs)
		 total+=count2
		 file3.Seek(0,0)

		 file3.Write([]byte{ count2}) //////类型

	 }
	 //临时文件创建，储存




}