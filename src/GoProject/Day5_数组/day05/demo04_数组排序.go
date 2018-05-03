package main

import "fmt"

func main() {
	/*
	排序：从小到大(升序)，或者从大到小(降序)
		让数组中的数据按照从小到大排列。
	 */
	arr1 := [5]int{15, 23, 8, 10, 7}
	fmt.Println(arr1) // 7 8 10 15 23

	for i:=1;i< len(arr1);i++{// i表示第几轮排序
		for j:=0;j<len(arr1) - i;j++ {
			if arr1[j] > arr1[j+1] {
				arr1[j], arr1[j+1] = arr1[j+1], arr1[j]
			}
		}
	}
	fmt.Println(arr1)


}
