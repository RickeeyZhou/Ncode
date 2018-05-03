package main

import (
	"os"
	"fmt"
	"strconv"
	"io"
)

func main()  {
	/*
	思路：断点续传，边复制，边记录复制的数据量
	 */
	srcName:="/home/rickeey/测试/aa.jpeg"
	destName:="/home/rickeey/测试/haha/aa.jpeg"
	file1,_:=os.Open(srcName)
	file2,_:=os.OpenFile(destName,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	tempName :=destName+"temp.txt" //C:\\Ruby\\bb.jpegtemp.txt
	file3,_:=os.OpenFile(tempName,os.O_RDWR|os.O_CREATE,os.ModePerm)
	defer file1.Close()
	defer file2.Close()
	//defer file3.Close()


	//1.从临时文件中读取已经存储的上次的拷贝的数据量
	totalBytes:=make([] byte,100)
	count1,_:=file3.Read(totalBytes) // 将已经拷贝的数据量读取到数组中
	totalStr := string(totalBytes[:count1])// 从数组中获取读取的数量量，-->string
	total,_ := strconv.Atoi(totalStr) //int
	fmt.Println("上次已经复制的数据量",total)

	//2.设置读写的位置
	file1.Seek(int64(total),0)
	file2.Seek(int64(total),0)
	dataBytes:=make([] byte, 1024)
	for{
		count2,err:= file1.Read(dataBytes)
		if err == io.EOF {
			fmt.Println("已经复制到文件末尾。。", total)
			file3.Close()
			os.Remove(tempName)
			break
		}

		file2.Write(dataBytes[:count2])
		total += count2
		file3.Seek(0,0)
		totalStr=strconv.Itoa(total)
		file3.WriteString(totalStr)

		if total > 30000{
			panic("意外断点了，，假装的。。。。")
		}
	}




}





