package main

import "fmt"

func main()  {
	/*
	  1.创建一个map用于存储一个人的信息：
	"name"：王二狗
	"age"：30
	sex：男性
	address："北京市xx路xx号"

2.打印遍历该map
3.修改sex为女性
	  4.如果想map中添加重复的key，会如何？
	   */
	   map1 := make(map[string]string) // 非nil
	   //map1[key]=value
	   map1["name"]="王二狗"
	   map1["age"]="30"
	   map1["sex"]="男性"
	   map1["address"] = "北京市猫眼胡同38号"
	   //for key,value := range  map1{
	   //	fmt.Println(key,"--->",value)
	   //}
	   for key,_:=range map1{
	   		fmt.Println(key, map1[key])
	   }

	   map1["sex"] = "女性"
	   fmt.Println(map1)

	   delete(map1, "address")
	   fmt.Println(map1)
	   map1["name"] = "三胖"

}
