package main

import "fmt"

func main()  {
	//1.变量
	a := 100
	//2.数组
	b := [4] int{1,2,3,4} // 一维数组
	fmt.Println(a)
	fmt.Println(b)
	//2.二维：一维的一维
	// 数组名:=[][]int{.}
	c := [3][4] int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12}}
	fmt.Println(c)
	fmt.Println("二维数组的长度：len(c)-->",len(c)) //二维数组中几个一维
	fmt.Println(c[0]) //第一个一维素组
	fmt.Println("一维数组的长度：len(c[1])-->",len(c[1])) //第二个一位数组的长度
	fmt.Println(c[0][0]) //第一个一维数组的第一个数据
	fmt.Println(c[1][3]) //第二个一维数组的第4个数据
	//fmt.Println(c[2][4]) // 数组的下标不能越界
	//，遍历该二维数组
	for i:=0;i<len(c);i++{ //控制第几个一维
		//i:0,第一个一维
		//i:1，第二个一维
		//i:2,第三个一维
		for j:=0;j<len(c[i]);j++{//打印该一维中第几个数的
			fmt.Print(c[i][j],"\t")
			//c,二维数组，c[i]，二维中第几个一维，c[i][j],一维中第几个数
		}
		fmt.Println()
	}
	//第一个一维
	//for j:=0;j<len(c[0]);j++{
	//	fmt.Print(c[0][j],"\t")
	//}
	//fmt.Println()
	//
	//for j:=0;j<len(c[1]);j++{
	//	fmt.Print(c[1][j],"\t")
	//}
	//fmt.Println()
	//
	//for j:=0;j<len(c[2]);j++{
	//	fmt.Print(c[2][j],"\t")
	//}
	//fmt.Println()
}
