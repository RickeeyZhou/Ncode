package main

import "fmt"

func main()  {
	/*
	map集合是引用类型：
	传递时候，传递的是地址。
	 */
	map1:=map[string]int{
		"吴琼":50000,
		"朱磊":80000,
		"王二狗":7800,
		"三胖":6000}
	fmt.Println(map1)
	map2:=map1
	fmt.Println(map2)
	map2["朱磊"] = 9000000
	fmt.Println(map2)
	fmt.Println(map1)

	// type
	fmt.Printf("%T", map1) // map[string]int

	//map3:= make(map[int]int) // map[int]int
	//map1 = map3
	//fmt.Print(map1,map3)

	/*创建一个map，存储以下数据，并按照key的大小顺序打印输出。
	1:"HR"
	2:"程序员"
	3:"程序员鼓励师"
	4:"行政专员"
	5:"前台妹子"
	*/
}
