package main

import (
	//"time"
	"fmt"
	"crypto/sha256"
)

//挖矿
func isvaliddifficlt (hash string,difficulty int )bool{
	b:=len(hash)
	var i int
	for i=0;i<b;i++{
		if hash[i]!='0'{
			break
		}
	}
	return i>=difficulty
}
func main(){
	count:=0
	hashstring:=""
	//timestam1:=time.StampNano
	// 挖矿
	// 区块数据: prehash string ; timestamp: " //" ; data string ; nonce string(nonce)
	//创建hash算法对象
	hash1:=sha256.New() ///*
	nonce:=0 // 参数
	Prehash:="0000e1343c8d9ec8a8996f0c1c8d1f9f1d954a750c3db5525205f401c516222d"
	for !isvaliddifficlt(hashstring,4){
		count++
		nonce++
		hashstring=string(nonce)+"Thu, 26 Apr 2018 03:30:54 GMT"+Prehash
		hash1.Write([]byte(hashstring))
		hashstring=fmt.Sprintf("%x",hash1.Sum(nil)) //每次生成的hash是切片数组,把它转换成十六进制字符串 字节数组
		fmt.Println(hashstring)
	}
	fmt.Println(count)
}
