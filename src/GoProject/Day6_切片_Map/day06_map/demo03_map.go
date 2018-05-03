package main

import "fmt"

func main()  {
	/*
	map：
		概念：存储一组无序的键值对的集合。
		语法：
			A：创建：如果一个map仅有声明，就是nil map，不能直接存储数据
			var map名 [key类型]value类型
			map名:=map[key类型]value类型{}
			map名:=make(map[key类型]value类型)

			B：向map中添加数据
				map[key] = value

			C：获取map的数据：根据key获取value，如果key 不存在
				value, ok := map[key]
					根据key获取对应的value
						如果key存在，value就是对应的数据，ok是true
						如果key不存在，value就是该类型的默认值，ok是false
				遍历：
			D：删除键值对
				delete(map,key)
				根据key删除键值对，如果key不存在，删除失败。
			E：修改：
				根据key修改
	 */
	 //1.创建map
	 var map1 map[int]string
	 map2:= map[int]string{1:"王者农药",2:"绝地求生",3:"连连看",4:"传奇霸业",5:"消消乐"}
	 map3:=make(map[int]string)
	 fmt.Println(map1)
	 fmt.Println(map2)
	 fmt.Println(map3)
	//2.nil map
	fmt.Println(map1 == nil, map2 ==nil, map3==nil)//true false false
	fmt.Printf("%p,%p,%p\n",map1,map2,map3)

	 //3.存储键值对，映射项
	 if map1 == nil{
	 	fmt.Println("map1上未创建。。")
	 	map1 = make(map[int]string)
	 }
	 map1[1] = "王二狗"
	 map1[2] = "李小花"
	 map1[3] = "三胖"
	 fmt.Println(map1)

	 // 4.获取map中的数据
	 fmt.Println(map1[2])
	 fmt.Println(map1[20]) // 根据key获取对应的value值，如果key不存在，获取到的是该值类型的默认值，int：0，string：""
	 //value,ok:=map[key]
	 val1, ok:= map1[20]
	if ok == true {
		fmt.Println(val1)
	}else {
		fmt.Println("您获取的key不存在。。。没有对应的value值")
	}
	/*
	遍历map，配合range

	range-->array/slice
		range:index,value
	range-->string
		range:index,char
	range-->map
		range:key,value
	  */

	  //4.删除某个键值对
	  delete(map1, 30)
	  fmt.Println(map1)
	  // 5.根据key修改value值
	  map1[3] = "如花"
	  fmt.Println(map1)


}
