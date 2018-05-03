package main

import "fmt"

func main()  {
	// 练习2：格式化输出以下数据
	a := 100
	b := 3.14
	c := "hello"
	d := `王二狗住在隔壁`
	e := true
	f := 'A' // int32
	//原型输出
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", e, e)
	fmt.Printf("%v, %T\n", f, f)

	fmt.Println("----------------")
	fmt.Printf("%d,\n",a)
	fmt.Printf("%f,\n",b)
	fmt.Printf("%s,\n",c)
	fmt.Printf("%s,\n",d)
	fmt.Printf("%t,\n",e)
	fmt.Printf("%q,\n",f)
}
