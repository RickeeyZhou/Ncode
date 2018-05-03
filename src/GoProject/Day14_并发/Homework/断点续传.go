package main

import (
	"os"
	"fmt"
	"io"
	"strconv"
)

func main() {
	srcname := "/home/rickeey/测试/aa.jpeg"
	desname := "/home/rickeey/测试/haha/aa.jpeg"

	src, _ := os.Open(srcname)
	des, _ := os.OpenFile(desname, os.O_RDWR|os.O_CREATE, 0777)
	Temp, _ := os.OpenFile("/home/rickeey/测试/temp.txt", os.O_CREATE|os.O_RDWR, 0777)

	//拷贝
	bs := make([]byte, 1024)
	total := make([]byte, 100)

	coutn1,_:=Temp.Read(total) //total []byte 100 : 21232
	totalstr := string(total[:coutn1])
	Total, _ := strconv.Atoi(totalstr) //  先读取临时文件的值
	fmt.Println("last time data is:",Total)

	src.Seek(int64(Total), 0)
	des.Seek(int64(Total), 0)

	for {
		count, err := src.Read(bs)
			if err != nil {
			fmt.Println(err)
			return
			}
		Total += count
			if err == io.EOF {
			des.Write(bs[:count])
			Temp.Close()
			break
			}
		des.Write(bs)
		Temp.Seek(0,0)
		signstring:=strconv.Itoa(Total)
		Temp.WriteString(signstring)

		if Total>3000{
			panic("人在中断")
		}
	}


}

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
/*
第一层 拷贝文件，读文件 ，写文件
第二层 遇到 程序中断时，如何从中断处恢复，涉及到记录中断点，此时用到临时文件记录
第三层 记录的是一个整数，并且会很大，但是读文件的功能是读取一组字节，你显示什么，
我读什么，计算机并不知道什么意义，就像字符。
但是 我记录的中断点，是一个整数，意义在于他的数值。所以类型就不一致了。
第四层 先将数值转换成字符串 ，要用时，在转换回来。
 */