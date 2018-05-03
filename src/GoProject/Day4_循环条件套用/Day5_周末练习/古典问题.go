package main

import "fmt"

func main(){       //斐波拉契数列 加 数组
	var v[13] int64
		v[1]= 1
		v[2]= 1
//声明 【12】个 但是是从0开始 n最多等于11
	for n:=3;n<=12;n++{
		v[n]=v[n-1]+v[n-2]
		fmt.Printf("%d月：%d\n",n,v[n])
	}
}