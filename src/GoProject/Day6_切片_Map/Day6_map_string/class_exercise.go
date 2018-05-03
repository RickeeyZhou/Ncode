package main

import "fmt"
import "strings"
func main() {
	s1 := "HelloWorld"
	s2 := "20180403课堂笔记.txt"
	fmt.Println(strings.Contains(s1, "ll"))
	fmt.Println(strings.ContainsAny(s2, "tt"))
	fmt.Println(strings.Index(s1, "llo"))
	fmt.Println(strings.LastIndex(s1,"."))

	if strings.HasPrefix(s2,"201804"){
		fmt.Println("This is a April note book")
	}
	if strings.HasSuffix(s2,"txt"){
		fmt.Println("This is a .txt file")
	}
	//
	fmt.Println(strings.Count(s1,"1"))
	//
	s3:="memeda,123a,hello,ddd"
	arr1:=strings.Split(s3,",")
	fmt.Println(arr1)
	fmt.Printf("%T\n",arr1)
	fmt.Println(len(arr1))

	arr2:=strings.SplitN(s3,",",-1)
	fmt.Println(arr2)
	s4:=strings.Join(arr2,",")
	fmt.Println(s4)

	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s1))

	s5:="****,he llo world*.*+++"
	fmt.Println(s5)
	fmt.Println(strings.Trim(s5,","))

	fmt.Println(strings.Replace(s1,"1","*",-1))

	fmt.Println(s1)
	slice6:=s1[:5]
	fmt.Println(slice6,"\n")

	fmt.Println(s1[0])
	s1="memeda"
	fmt.Println(s1)


}