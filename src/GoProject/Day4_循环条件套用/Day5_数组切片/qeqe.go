package main

import "fmt"

func main()  {
	const  n = 10
	index:=0//遍历数组
	count:=0//计数器
	sum:=0//统计退出的人数
	arr := [n] bool{}
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		arr[i] = true
	}
	fmt.Println(arr)
	for sum < n - 1{ //sum :0
		if arr[index] == true{ // index:0,1,2,3
			count++ // 1,2,3
			if count == 3{
				count = 0
				sum++ // 1,2,3
				arr[index] = false
			}
		}
		index++ // 1,2,3
		if index == n{
			index = 0
		}
	}
	//输出结果
	for i:=0;i<len(arr);i++{
		if arr[i]==true{
			fmt.Println(i+1)
		}
	}
	fmt.Println(arr)
}
