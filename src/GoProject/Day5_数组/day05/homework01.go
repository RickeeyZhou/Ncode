package main

import "fmt"

func main()  {
	/*
	Go是强类型语言，也叫静态语言。运算符两端的操作数的类型必须一致
	=，+，-，*，/....
	Type(value)-->T(v)
	 */
	 for i:=380;i<=450;i++{//i表示总人数
		for j:=1;j<i;j++{//j表示男生人数
			k:=i-j
			if float64(i) * 76 == float64(j)*75+float64(k)*80.1{
				fmt.Println("男生人数：",j,"女生人数：",k)
			}
		}
	 }
}
