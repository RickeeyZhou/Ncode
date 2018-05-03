package main

import "fmt"

type Person2 struct {
	name string
	age int
}

func (p Person2) eat()  {
	fmt.Println("吃窝窝头。。。")
}

func (p Person2) getInfo() string {
	/*
	fmt包下：Printf()，格式化输出
	Springf(),格式化获取字符串
	 */
	msg := fmt.Sprintf("姓名：%s,年龄：%d", p.name,p.age)
	//fmt.Println(msg)
	return msg
}

type Student2 struct {
	Person2
	school string
}
//可以重写getInfo()
func (s Student2) getInfo() string{
	msg := fmt.Sprintf("姓名：%s，年龄：%d，学校：%s", s.name,s.age,s.school)
	return msg
}

func (s Student2) study()  {
	fmt.Println(s.name,",废寝忘食，头悬梁，锥刺股。自己相信就好。")
}

func main()  {
	p1 := Person2{"王二狗",30}
	fmt.Println(p1.getInfo())// 父类对象，调用父类的方法

	s1:= Student2{Person2{"李小花",18},"蓝翔技校"}
	fmt.Println(s1.getInfo())//子类重写的

	s1.eat() // 子类访问父类的方法
	s1.study()// 调用子类新增的方法
}
