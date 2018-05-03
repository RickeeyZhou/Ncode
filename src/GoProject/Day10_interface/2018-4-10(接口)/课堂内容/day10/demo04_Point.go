package main

import (
	"fmt"
	"math"
)

/*
定义一个类：Point
	x，y，z轴
创建点的对象：
	两个点

定义函数，方法，求两点之间的距离
 */

type Point struct {
	x float64//x轴的坐标
	y float64//y轴的坐标
	z float64//z轴的坐标
}

func (p Point) printInfo()  {
	fmt.Printf("x轴坐标：%.2f,y轴坐标：%.2f,z轴坐标：%.2f\n",p.x, p.y,p.z)
}
// 方法：求p到指定点p2的距离
func (p Point) getDistance3(p2 Point) float64 {
	dis := math.Sqrt((p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y)+(p.z-p2.z)*(p.z-p2.z))
	return dis
}

// 函数
func getDistance(x1,x2,y1,y2,z1,z2 float64) float64{
	dis := math.Sqrt((x1-x2)*(x1-x2) +(y1-y2)*(y1-y2)+(z1-z2)*(z1-z2))
	return dis
}
// 函数
func getDistance2( p1, p2 Point) float64  {
	dis := math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) +(p1.y-p2.y)*(p1.y-p2.y)+(p1.z-p2.z)*(p1.z-p2.z))
	return dis
}

func main()  {
	p1 := Point{0,0,0}
	p1.printInfo()

	p2 := Point{2.4,4.1,6.3}
	p2.printInfo()

	p3 := Point{1.4,6.8,3.7}

	res1 := getDistance(p1.x,p2.x,p1.y,p2.y,p1.z,p2.z)
	fmt.Println(res1)

	res2 := getDistance2(p1, p2)
	fmt.Println(res2)

	res3 := p1.getDistance3(p2)
	fmt.Println(res3)
	res4:=p1.getDistance3(p3)
	res5:=p3.getDistance3(p2)
	res6:=p2.getDistance3(p3)

	fmt.Println(res4)
	fmt.Println(res5)
	fmt.Println(res6)

}

