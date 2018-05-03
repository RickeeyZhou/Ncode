package main

import (
	"fmt"
	"strings"
)

func main()  {
	str:="aekjffjkJDJ294384848DKFJFJkdjfhfh2943845593nfnJRIEIFJ"
	count1:= 0//大写字母个数
	count2:=0//小写字母个数
	count3:=0//非字母
	for _,char:=range str{
		if char >='A' && char <= 'Z'{
			count1++
		}else if char>='a' && char <='z'{
			count2++
		}else{
			count3++
		}
	}
	fmt.Println("大写字母个数：",count1,"小写字母个数：",count2,"非字母个数：",count3)

	// 方法二：
	count1=0
	count2=0
	count3=0
	s1:="QWERTYUIOPASDFGHJKLZXCVBNM"
	s2:="qwertyuiopasdfghjklzxcvbnm"
	arr:= strings.Split(str,"")
	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		//Count, Index(),Contains
		//if strings.Index(s1, arr[i])!=-1{
		//	count1++
		//}
		if strings.Contains(s1, arr[i]){
			count1++
		}else if strings.Contains(s2,arr[i]){
			count2++
		}else{
			count3++
		}
	}
	fmt.Println("大写字母个数：",count1,"小写字母个数：",count2,"非字母个数：",count3)
	//方法三：

}
