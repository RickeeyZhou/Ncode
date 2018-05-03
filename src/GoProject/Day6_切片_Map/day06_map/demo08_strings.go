package main

import (
	"fmt"
	"strings"
)
func main()  {
	/*
	ctrl+q,查看一个函数方法的声明
	按住ctrl，点击函数方法名，查看源代码
	查API


	strings包下的关于字符串 的常用方法：
	Contains(s,substr),是否包含指定内容
	ContainsAny(s,chars),是否包含指定内容中的任意一个
	HasPrefix(),判断以指定内容开头
	HasSuffix(),判断以指定内容结尾
	Index(),搜索指定的内容，第一次出现的索引下标，如果没有返回-1
	LastIndex(),最后一次出现的索引内容，如果没有也是-1

	Count(),统计指定的内容出现的次数


	Split(s1,"")-->[]string
	SplitN(s1,"",n)-->[]string,指定次数的切割，不超过n次。如果全切n=-1
	Join([]strings,"")-->string,合并

	ToLower()
	ToUpper()

	Trim(),去除首尾的指定内容

	Replace(s,old,new,n),n替换的次数，-1表示全部替换
	Repeat(),重复的次数
	 */
	 s1:="HelloWorld"
	 s2:="20180403课堂笔记.txt"
	 //是否包含
	 fmt.Println(strings.Contains(s1,"O"))
	 fmt.Println(strings.ContainsAny(s1,"abcd"))
	 fmt.Println(strings.Index(s1,"llo")) //查找指定的内容，返回第一次出现的下标，如果没有返回-1。
	 fmt.Println(strings.LastIndex(s1,"l")) //返回最后一次出现的位置。
	 //判断以xx开头，结尾
	if strings.HasPrefix(s2,"201804") {
		fmt.Println("这个是4月份的笔记。。")
	}
	if strings.HasSuffix(s2,".txt"){
		fmt.Println("这个是一个文本文档。。")
	}
	 //统计
	 fmt.Println(strings.Count(s1,"l"))

	 //切割：
	 s3:="memed,123a,hello,ddd"
	 arr1:=strings.Split(s3,",")
	 fmt.Printf("%T\n",arr1)//[]string
	 fmt.Println(arr1)
	 fmt.Println(len(arr1))

	 arr2:=strings.SplitN(s3,",",-1)
	 fmt.Println(arr2)
	 s4:=strings.Join(arr2,",")
	 fmt.Println(s4)

	 fmt.Println(strings.ToLower(s1))//转小写
	 fmt.Println(strings.ToUpper(s1))//转大写


	 s5:="***hel llo world**+++"
	 fmt.Println(s5)
	 fmt.Println(strings.Trim(s5,"*+"))


	 fmt.Println(strings.Replace(s1,"l","*",-1))

	 fmt.Println(strings.Repeat(s1,3))

	 //字符串的截取-->
	 //strings.substring(start,end)
	 fmt.Println(s1)
	 s6:=s1[:5]
	 fmt.Println(s6)
	 //字符串的内容不可以改变
	 fmt.Println(s1[0])
	 s1 = "memeda"
	 fmt.Println(s1)


}
