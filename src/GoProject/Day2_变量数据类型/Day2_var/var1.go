package main

import "fmt"

func main() {
	//整型（数值型）
	var x int = 12
	var x2 = 12
	var x4 int
	x3 := 12.2
	Xx := int(x3)
	x3 = x3
	var (
		x5 = 5
		x6 = 1
	)
	var y1, y2 = 12, 12
	fmt.Printf("数值：%d 类型是%T\n", x, x)
	fmt.Printf("数值：%d 类型是%T\n", x2, x2)
	fmt.Printf("数值：%d 类型是%T\n", Xx, Xx)
	fmt.Printf("数值：%d 类型是%T\n", x4, x4)
	fmt.Printf("数值：%d 类型是%T\n", x5, x5)
	fmt.Printf("数值：%d 类型是%T\n", x6, x6)
	fmt.Printf("数值：%d 类型是%T\n", y1, y1)
	fmt.Printf("数值：%d 类型是%T\n", y2, y2)
	//浮点型

	//字符串
	var s1, s2, s3 string = "fuck", "you", "dick"
	var s4 = `fuck u again`
	fmt.Printf("%s %q %s;%T\n", s1, s2, s3, s1)
	fmt.Printf(s4)

	//布尔型
	t1, t2 := false, true
	t3 := (t1 == t2)
	fmt.Println(t1, t2)
	fmt.Printf("%t ,%t", t3, t3)

	//常量
	const c1, c2 int = 12, 13
	const c3 = 11.2
	const (
		v5 = 's'
		v4 = "fuck"
		v6 = true
	)
	//运算
	var sum, de, ch, sha, yu int
	sum = x + x5
	de = x - x5
	ch = x * x5
	sha = x / x5
	yu = x % x5
	fmt.Println(sum, de, ch, sha, yu)

	//验证变量声明类型和赋值类型不一样的情况
	var v1 int = 10 % 3
	fmt.Printf("\n value=%v type=%T", v1, v1)
	//直接赋值3.3会报错，体现了强语言的性质
	//但是如果右边赋值部分是一个表达式则会自动转换成整型顺利运行
	//解决：10/3 是商 商 =3
	b1, b2 := 31, 12
	fmt.Println()
	fmt.Println(b1 > b2, b1 == b2, b1 != b2)
	//a := 4
	//	b := 3
	//	res1 := a < b && b / 2 == 1 && a % 3 != 0  fals

	//	res2 := (a+b)*3 < a<<2 || (a-b) >0    ture 先运算 位运算 和四则运算

	//交换2个变量的值
	a := 3
	b := 2
	//人类思维 利用参数
	var c int
	c=a
	a=b
	b=c
	//位运算 运用二进制的逻辑关系
	a=a&b
	b=a|b
	//利用四则运算
	a--
	b++
	//利用


}
