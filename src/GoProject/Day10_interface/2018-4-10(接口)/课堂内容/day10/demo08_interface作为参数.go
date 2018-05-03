package main

import (
	"math"
	"fmt"
)

type Shape interface {
	area() float64
	peri() float64
}
type Circle struct {
	radius float64//半径
}

func (c Circle) area() float64  {
	return c.radius * c.radius * 3.14
}

func (c Circle) peri()float64  {
	return 2 * 3.14 * c.radius
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (t Triangle) area()float64  {
	p := t.peri() / 2
	area := math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
	return area
}
func (t Triangle) peri() float64  {
	return t.a + t.b + t.c
}

func testShape(s Shape)  {
	fmt.Println("面积是：",s.area(),"，周长是：",s.peri())
}
/*
定义一个接口：形状
	定义两个方法：
		周长()
		面积()

定义三个结构体分别实现该接口：
	三角形：周长()，面积()
		海伦公式：
	矩形：周长()，面积()
	圆形：周长()，面积()

在主函数中：分别创建三种形状的对象，并调用对应的周长和面积的方法

增加一个测试的的函数：
	testArea(接口形状)-->三角形，原型，矩形
	testPeri()-->三角形，圆形，矩形
 */

func main()  {
	t1 := Triangle{3,4,5}
	testShape(t1)
	testShape(Triangle{5,6,8})
	testShape(Circle{5})
}
