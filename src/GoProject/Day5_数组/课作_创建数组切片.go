package main

import "fmt"

func main() {
	var a [4] int
	a = [4]int{1, 2, 23, 3}
	a[3] = 2
	fmt.Print(a)
	var b = [4]int{2, 2, 3, 3}
	//我总喜欢用c语言的声明格式
	//数组在go是数据类型了，不再是对象
	var ab = [12]int{221, 12, 3, 12}
	a4 := [...]int{23, 23, 1, 23}
	a6 := [...]int{0: 22, 2}
	fmt.Print(len(a6), a4, ab, b, a)
	var fl=[2]float64{2.2,1.1}
	fmt.Print("\n",fl)
	st:=[3]string{"23232,","dsdd,","sds"}
	fmt.Print(st)
	var arr1=[10]int{5,4,3,7,10,2,9,8,6,1}
	sum:=0
	for i:=0;i<len(arr1);i++{
		sum+=arr1[i]
	}
	fmt.Println(sum)

}
