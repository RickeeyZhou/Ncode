package main

import "fmt"

func main()  {
	str2:="abadebd"
	arr4:=[]byte(str2)
	fmt.Println(arr4)
	slice2:=make([] byte,0,10) //[a,b,d]
	fmt.Println(slice2)
	for i:=0;i<len(arr4);i++{
		flag := true
		for j:=0;j<i;j++{
			if slice2[j] == arr4[i]{
				flag= false
				break
			}
		}
		if flag {
			slice2 = append(slice2, arr4[i])
			fmt.Println(slice2)
		}
	}
	fmt.Println(slice2)
}
