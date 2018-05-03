package main

import (
	"crypto/sha256"
	"time"
	"fmt"
)


func isvalidhashdifficlut (  h  string , difficlit int )bool{
// 判断哈希值是否符合难度条件
	var i int
	for i=0;i<len(h);i++{
		if string(h[i])!="0"{
			break
		}
	}
	return i>=difficlit
}
func main(){
	h:=""
	var nonce int
	timeStamp:=time.StampNano //string
	shaobj:=sha256.New()
	//hashstring= shaobj.Write()
	prestring:="e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"  // 前哈希
	for ;!isvalidhashdifficlut(h,4);{
		nonce++
		input :=prestring+h+timeStamp+string(nonce)
		shaobj.Write([]byte(input)) /// 生成哈希值
		h=fmt.Sprintf("%x",shaobj.Sum(nil)) // []byte
		fmt.Println(h)
	}


}