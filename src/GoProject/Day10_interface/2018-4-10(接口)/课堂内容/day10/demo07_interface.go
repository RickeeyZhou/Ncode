package main

import "fmt"

//定义一个接口
type USB interface {
	//只有功能的描述：就是方法的声明，没有实现
	name() string //描述一个功能：该功能可以获取名称
	plugIn() // 描述一个功能：能够插入工作
}
//结构体：--->实现USB接口
type Mouse struct {
	m_name string // 鼠标名称
}

func (m Mouse) name()string  {
	return m.m_name
}
func (m Mouse) plugIn()  {
	fmt.Println("USB设备：",m.name(),"连入电脑开始工作。。。")
}

type FlashDisk struct {
	f_name string//U盘名称
}

func (f FlashDisk) name() string {
	return f.f_name
}

func (f FlashDisk)plugIn()  {
	fmt.Printf("USB设备：%s，连入主机开始准备工作。。",f.name())
}

func main()  {
	/*
	接口：interface
		接口中：约定了行为功能，但是没有具体的实现。
			也不能有字段的定义

		接口中：一个或多个方法的声明。

	只要能够实现该接口中的所有的方法，那么就叫做该接口的实现类。
	 */

	 m1:=Mouse{"罗技红色"}
	 fmt.Println(m1.name())
	 m1.plugIn()

	 f1:=FlashDisk{"闪迪32G"}
	 f1.plugIn()
}

