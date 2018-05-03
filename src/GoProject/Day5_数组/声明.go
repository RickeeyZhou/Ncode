package main

import "fmt"

func main() {
	arr1 := [5]int{15, 23, 8, 10, 7}
	fmt.Print(arr1)

	for j := 0; j < 4; j++ {
		if arr1[j] > arr1[j+1] {
			arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
		}
	}
	fmt.Print(arr1)

	for j := 0; j < 3; j++ {
		if arr1[j] > arr1[j+1] {
			arr1[j], arr1[j+1] = arr1[j+1], arr1[j]

		}
	}
	fmt.Println(arr1)
	for j := 0; j < 2; j++ {
		if arr1[j] > arr1[j+1] {
			arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
		}
	}
	fmt.Println(arr1)
	for j := 0; j < 1; j++ {
		if arr1[j] > arr1[j+1] {
			arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
		}
	}
	fmt.Print("\n", arr1)
	//var arr2[5] int
	arr2 := [5]int{23, 12, 44, 2, 7}
	for a := len(arr2); a > 1; a-- {
		for b := 0; b < a-1; b++ {
			if arr2[b] > arr2[b+1] {
				arr2[b], arr2[b+1] = arr2[b+1], arr2[b]
			}
		}
	}
	fmt.Print(arr2)

	c := [3][4] int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}
	fmt.Println(c)
	fmt.Println("二维数组的长度：len(c)-->", len(c))       //二维数组中几个一维
	fmt.Println(c[0])                              //第一个一维素组
	fmt.Println("一维数组的长度：len(c[1])-->", len(c[1])) //第二个一位数组的长度
	fmt.Println(c[0][0])                           //第一个一维数组的第一个数据
	fmt.Println(c[1][3])                           //第二个一维数组的第4个数据
	//fmt.Println(c[2][4]) // 数组的下标不能越界
	//，遍历该二维数组
	for j := 0; j < len(c[0]); j++ {
		fmt.Print(c[0][j], "\t")
	}
	for j := 0; j < len(c[1]); j++ {
		fmt.Print(c[1][j],"\t")
	}
	for j := 0; j < len(c[2]); j++ {
		fmt.Print(c[2][j],"\t")
	}
	fmt.Print("\n")
	for i:=0;i<len(c);i++{
		for j:=0;j<len(c[i]);j++{
			fmt.Print(c[i][j],"\t")
		}
	}

}
