package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	/*
	readDir(dirname)-->[] FileInfo
		获取对应目录下的内容，只一层。
	 */
	 dirName:="C:\\Ruby\\pro"
	 fileInfos,_:=ioutil.ReadDir(dirName)
	 fmt.Println(len(fileInfos))
	 for i,fi:= range fileInfos{
	 	fmt.Printf("第%d个子内容：\n",i+1)
	 	fmt.Println("\t名字：",fi.Name())
	 	fmt.Println("\t是否是目录：",fi.IsDir())
	 }
}
