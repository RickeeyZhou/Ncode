package main
/*
整数除法 取后面精确小数
小数部分 取决于余数
所以[0,1]的任意浮点数可以表示任意两个整数的小数部分
算法 一个位 一个位的取值
a*10）% b 的模  不管是整数 还是小数 部分，理论上都可以用该算法完成
但是无法确定 小数点取在哪
 */
import "fmt"

func main() {
	var a, b int
	fmt.Scanf(" %d / %d ",&a,&b)  ///在格式化输入中
	fmt.Println("你想精确小数点后几位？")
	var n int
	fmt.Scanln(&n)
	var Int int
	var Float int
	Int = a/b
	fmt.Print(Int,".")
	Float=a%b
	for temp:=0 ;n>0;n--{
		temp=(Float*10)/b
		fmt.Print(temp)
		Float=(Float*10)%b
	}
}