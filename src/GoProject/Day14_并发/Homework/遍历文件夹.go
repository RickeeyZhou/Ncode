package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	dirname := "/home/rickeey/测试"
listFiles(dirname)
}
func listFiles(dirname string) {
	fileInfo, _ := ioutil.ReadDir(dirname)
	fmt.Println()
fmt.Println(fileInfo)
	fmt.Println()
	for _,fi:=range fileInfo{
		filename:=dirname+"/"+fi.Name()
		fmt.Println(filename)
		if fi.IsDir(){
			listFiles(filename)
		}

	}
}
