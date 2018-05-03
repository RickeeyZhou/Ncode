package main

import "fmt"

func main()  {
	/*
	可变参数：
		如果一个函数的参数，类型一致，但是个数不定，那么可以使用可变参数
	语法格式：
		参数名 ... 类型
	数量不定：从0到多个
	注意点：
		1.一个函数最多只能有一个可变参数
		2.如果参数列表中，还有其他的参数，那么可变参写在参数列表的最后
	 */
	 fmt.Println(getSum5(1,2,3,4,5))
	 s1:=[] int{1,2,3,4,5,6,7,8,9,10}
	 fmt.Println(getSum5(s1...))

	 s1 = append(s1, 100,200,300)
	 s2 := []int{1,3,5,7}
	 s1 = append(s1,s2...)
}
// 定义一个函数，用于求n个整数的和
func getSum5( nums ... int) int { // []int s1 string,s2 string,
	fmt.Printf("%T,%d\n", nums,len(nums))
	sum:=0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	return sum
}
