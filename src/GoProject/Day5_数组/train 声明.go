package main

import "fmt"
func main(){
	a :=  [2]string{"22","11"}
	fmt.Printf("%T,%v,%T,%v ",a,a,a[0],a[0])

	a[0],a[1]=a[1],a[0]
	 b:=[...]string{"",""}
	a ,b = b ,a //交换值
	var f[2] int
	fmt.Print(f)

	s1:=[]int{1,2,3,4,5,6}
	fmt.Println(s1)
	fmt.Print(len(s1),cap(s1))
	fmt.Printf("%T\n",s1)

	s2:=[]int{}
	fmt.Print(len(s2),cap(s2))

	s3:=make([]int,3,8)
	fmt.Println(s3)
	fmt.Println(len(s3),cap(s3))

	b1:=[10]int{1,2,3,4}
	s4:=b1[1:3]
	fmt.Print(s4,s2)
}