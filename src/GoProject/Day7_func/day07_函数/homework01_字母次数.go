package main

import (
	"fmt"
	"strings"
	"sort"
)

func main()  {
	/*
	ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN,
	每个字母的次数
	 */
	 str:="ECALIYHWEQAEFSZCVZTWEYXCPIURVCSWTDBCIOYXGTEGDTUMJHUMBJKHFGUKNKN"
	 arr:=strings.Split(str,"")
	 fmt.Println(arr)
	 sort.Strings(arr)
	 fmt.Println(arr)

	 for i:=0;i<len(arr);{
		 count:=0
		 for j:=i;j<len(arr);j++{
			 if arr[j] == arr[i]{
				 count++
			 }else {
			 	break
			 }
		 }
		 fmt.Printf("%s,次数是：%d\n",arr[i],count)
		 i += count
	 }

	 fmt.Println("方法二。。。。")
	 map1:= make(map[byte]int)
	 for i:=0;i<len(str);i++{
	 	//1.根据每个字母，获取map中对应的value
	 	value, ok := map1[str[i]]
	 	//2.如果ok为true，map中已成存储该字母，value加1
	 	if ok == true{
	 		value++
		}else{
			value = 1// 设置value就是1.
		}
		//3.更新map
		map1[str[i]] = value
	 }
	 s1:=make([]int,0,len(map1))
	 for key,_:= range map1{
	 	//fmt.Printf("%c,%d\n", key, value)
	 	s1=append(s1, int(key)) // byte-->int
	 }
	 sort.Ints(s1)
	 for _,k:= range s1{
	 	fmt.Printf("%c,%d\n", k,map1[byte(k)])
	 }

	 /*
	 A-Z:65-90
	 a-z:97-122
	  */
	for i:=65;i<=90;i++{
		num:=strings.Count(str,string(i))
		fmt.Printf("%q的个数%d\n",i,num)
	}
	/*
		数组：index，value
		A-Z,a-z 65 122
		str:="aekjffjkJDJ294384848DKFJFJkdjfhfh2943845593nfnJRIEIFJ"
		0-122
		123个0

		arr2[97] = 1
		 */
	arr2:=[123]byte{}
	for _,char:=range str{
		arr2[char]++
	}
	fmt.Println(arr2)
	for i,v:=range arr2{
		if v == 0{
			continue
		}
		fmt.Printf("%c,%d\n",i,v)
	}

	ss1:="abcd" // string
	ss2:=ss1[0:2] // string,substring()
	fmt.Printf("%T,%T\n",ss1,ss2)
	//substring(start,end)

	//substr(start,count)

	//sub(),sup()

	arr3:=[4]int{1,2,3,4} //[4]int
	slice1:=arr3[0:2]// []int
	fmt.Printf("%T,%T\n",arr3,slice1)



}
