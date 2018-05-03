package main

import (
	"fmt"
	"os"
)

func main()  {
	/*
	文件操作：
	 */
	 fileInfo, err:=os.Stat("C:\\Ruby\\pro\\aa.txt")
	 fmt.Println(err) //<nil>
	 fmt.Printf("%T\n",fileInfo) //*os.fileStat
	if err != nil{
		fmt.Println("文件操作有误，",err)
		return
	}
	// 否则查看文件信息
	fmt.Println(fileInfo.Name()) //aa.txt
	fmt.Println(fileInfo.IsDir()) //判断是否是目录，false
	fmt.Println(fileInfo.Size()) // 10,文件大小
	fmt.Println(fileInfo.ModTime())// 修改时间
	fmt.Println(fileInfo.Mode()) //-rw-rw-rw-   0666




}
