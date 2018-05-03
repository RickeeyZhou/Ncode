package main

import (
	"io/ioutil"
	"fmt"
)

func main()  {
	dirname:="C:\\Ruby\\pro"
	//1层：mmm，3个子内容
	//fileInfos,_:=ioutil.ReadDir(dirname)
	//fmt.Println(len(fileInfos))
	//for _,fi := range fileInfos{
	//	fmt.Println(fi.Name(),fi.IsDir())//C:\\Ruby\\pro\\mmm\\nn
	//	if fi.IsDir(){
	//		fileInfos2,_:=ioutil.ReadDir(dirname+"\\"+fi.Name())
	//		fmt.Println(len(fileInfos2))
	//		for _,fi2:=range fileInfos2{
	//			fmt.Println("\t",fi2.Name(),fi2.IsDir())
	//		}
	//	}
	//}
	listFiles(dirname)

}

func listFiles(dirname string)  {
	fileInfos,_:=ioutil.ReadDir(dirname)
	for _,fi:=range fileInfos{
		filename := dirname+"\\"+fi.Name()
		fmt.Println(filename)
		if fi.IsDir(){
			//继续遍历fi这个目录
			listFiles(filename)
		}
	}
}

