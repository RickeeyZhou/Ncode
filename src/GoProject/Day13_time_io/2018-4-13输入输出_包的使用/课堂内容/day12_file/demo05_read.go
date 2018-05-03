package main

import (
	"os"
	"fmt"
	"io"
)

func main()  {
	file,err:= os.Open("C:\\Ruby\\pro\\aa.txt")
	if err!=nil{
		fmt.Println("文件打开失败。。")
		return
	}
	defer file.Close()
	//读取数据
	bs := make([] byte,4)
	count := 0//每次读取的数量
	for {
		count, err = file.Read(bs)
		if err ==io.EOF{
			fmt.Println()
			fmt.Println("数据读取完毕。。")
			break
		}
		fmt.Print(string(bs[:count]))
	}


	/*
	//第一次读取,通过file读取对应的文件，将数据读入到bs中，返回值是本次实际读入的数量
	count ,err = file.Read(bs)
	fmt.Println(err) //<nil>
	fmt.Println(count) //4
	fmt.Println(bs) //[97 98 99 100]
	fmt.Println(string(bs)) //abcd

	//第二次读取
	count,err = file.Read(bs)
	fmt.Println(err) //<nil>
	fmt.Println(count) //4
	fmt.Println(bs) //[101 102 103 104]
	fmt.Println(string(bs)) //efgh

	//第三次读，读取2个
	count,err = file.Read(bs)
	fmt.Println(err) //<nil>
	fmt.Println(count) //2
	fmt.Println(bs) //[101 102 103 104]
	//fmt.Println(string(bs)) //efgh  ijgh
	fmt.Println(string(bs[:count])) //ij
	//第四次读取,读取到文件末尾时，返回error为io.EOF,，而且count的值为0
	count, err= file.Read(bs)
	fmt.Println(err) //EOF
	fmt.Println(count) //0
	*/
	//练习1：熟悉Read(),
	//练习2：修改以上代码为循环，循环的退出条件。
}
