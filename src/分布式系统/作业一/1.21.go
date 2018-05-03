package main

import "fmt"

type student struct {
	Name string
	Age int
}
func pase_student() map[string]student{
	m:=make(map[string]student)
	stus:=[]student{
		{Name:"zhou",Age:24},
		{Name:"Li",Age:23},
		{Name:"Wang",Age:22},
	}
	for _,stu:=range stus{
		m[stu.Name]=stu
		fmt.Println(&stu)
	}
	return m
}
func main(){
	students:=pase_student()
	fmt.Println(students)
	for k,v:=range students{
		fmt.Printf("key=%s,value=%v\n",k,v)
	}
}
/*
pase_student的函数 创建了 学生类型的切片,里面存了三个学生对象.该切片Stus
然后遍历了切片,将学生对象逐个放入map映射中,
最后将map 返回给主函数的studens
故 students的数据类型也是 map
程序的最后遍历 map
将map进行了无序输出,结果理论上会出现6中情况
打印students 发现map[zhou:0xc4200b0000 Li:0xc4200b0000 Wang:0xc4200b0000]
储存的地址是一样的,说明当时创建键值对时,
是因为遍历切片 stus时, 储存只用到的stu 都是同一个,
改进方法,将&stu-->stu
改进后的结果:
&{zhou 24}
&{Li 23}
&{Wang 22}
map[zhou:{zhou 24} Li:{Li 23} Wang:{Wang 22}]
key=Li,value={Li 23}
key=Wang,value={Wang 22}
key=zhou,value={zhou 24}

Process finished with exit code 0
 */