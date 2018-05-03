package main

import "fmt"

func main() {
	s1:=[]int{}
	s2:=[]int{}
	arr:=[]int{}
	fmt.Println(len(s1),cap(s1))
	s1=append(s1,10)
	fmt.Println(s1)
	fmt.Println(len(s1),cap(s1))
	s2=append(s1,21,21)
	fmt.Println(s2)
	for i:=1;i<=10;i++{
		s1 = append(s1,i)
		fmt.Println("lenth:", len(s1), "capii:", len(s2))
	}
		s2=arr[:5]
		s3:=arr[2:7]
		fmt.Println(arr)
		fmt.Println(s2)
		fmt.Println(s3)
		s2[3]=100
		fmt.Println(s2)
		fmt.Println(s3)
		fmt.Println(arr)
		s2=append(s2,1,21,1,1,1,)
		s4:=[]int{1,2,23,3}
		s5:=make ([]int,2,)
		for i:=0;i<len(s5);i++{
			s5[i]=i*2+1
		}
		s4=append(s4,s5...)
		fmt.Pringln(s4)

		index:=3
		s4=append(s4[:index],s4[index+1:]...)
		fmt.Println(s4)
	}

