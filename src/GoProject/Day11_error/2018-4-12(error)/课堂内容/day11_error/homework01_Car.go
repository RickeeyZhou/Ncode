package main

import "fmt"

//1.定义接口表示发动机
type IEngine interface {
	start()
	speedUp()
	stop()
}
//2.定义结构体，实现接口
type YAMAHA struct {
	name string
}
//3.实现方法
func (y YAMAHA) start()  {
	//启动：60，加速80，停止0
	fmt.Println(y.name,"，启动，速度60")
}
func (y YAMAHA) speedUp()  {
	fmt.Println(y.name,"，加速，速度80")
}
func (y YAMAHA) stop()  {
	fmt.Println(y.name,"，停止，速度0")
}
type HONDA struct {
	name string
	speed int
}

func (h *HONDA) start()  {
	//启动：40，加速120，停止0
	h.speed = 40
	fmt.Println(h.name,"，启动，速度,",h.speed)
}
func (h *HONDA) speedUp()  {
	h.speed += 80
	fmt.Println(h.name,"，加速，速度",h.speed)
}
func (h *HONDA) stop()  {
	h.speed = 0
	fmt.Println(h.name,"，停止，速度,",h.speed)
}

type Car struct {
	iEngine IEngine //车里有个发送机,接口类型
}
func (c Car) testIEngine(){
	//测试发动机
	c.iEngine.start()
	c.iEngine.speedUp()
	c.iEngine.stop()
}
func main()  {
	c1:=Car{}
	//c1.iEngine = YAMAHA{"YAMHA"}
	i:=HONDA{"HONDA",0}
	c1.iEngine = &i
	c1.testIEngine()

	//c2 := Car{HONDA{"HONDA"}}
	//c2.testIEngine()

}
