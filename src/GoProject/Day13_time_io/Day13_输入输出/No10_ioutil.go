package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

func main(){
	fileName:="/home/rickeey/测试/aa.jpeg"
	bs,err:=ioutil.ReadFile(fileName)
	fmt.Println(err)
	fmt.Println(bs)
	fmt.Println(string(bs))

	fileName2:="/home/rickeey/测试/aa.txt"
	s1:="fucking the gffffirl ...."
	os.OpenFile("/home/rickeey/测试/aa.txt",os.O_CREATE|os.O_RDWR,0777)
	err1:=ioutil.WriteFile(fileName2,[]byte(s1),0)
	fmt.Println(err1)

	s2:="weeqeqweqweqweqeqweq"
	reader1:=strings.NewReader(s2)
	fmt.Println(reader1)
	bs2,_:=ioutil.ReadAll(reader1)
	fmt.Println(string(bs2))
	fmt.Printf("%T",bs2)


}
