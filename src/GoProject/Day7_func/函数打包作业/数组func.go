package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

func main(){
	var a[5] int=[5]int {1,2,3,4,5}
	var arr[3]string
	arr=[3]string  {"23","adad","12312"}
	c:=[4]float32{1.212,1212.2,122,123}
	for index,value:=range a {
		fmt.Printf("a[%d]=%d",index,value )
	}
	for index,value :=range arr{
		fmt.Printf("arr[%d]=%s",index,value)
	}
	for index,value :=range c{
		fmt.Printf("arr[%d]=%f",index,value)
	}

	var a2 int
	fmt.Scanln(&a2)
	fmt.Println(code(a2))
	fmt.Println()
	aa4:=randomarr()
	fmt.Println(aa4)
	sli:=aa4[:]
	fmt.Println(rankarr(sli))

	//////////////////////////////////////
	fmt.Println(sli2())
}
func code(a int  )[4]int{
	arr:=[4]int{}
	arr[0]=a/1000
	arr[1]=a%1000/100
	arr[2]=a%1000%100/10
	arr[3]=a%1000%100%10
	for index,value:=range arr{
		arr[index]=(value+5)%10
	}
	arr[0],arr[3]=arr[3],arr[0]
	return arr
}
func randomarr()[10]int {
	arr:=[10]int{}
	rand.Seed(time.Now().Unix())
	for index ,_:=range arr {
		arr[index]=rand.Intn(100)+1
	}
	return arr
}
func rankarr( sli[]int)[]int {
	sort.Ints(sli)
	return sli
}
func sli2()[]int {
sli:=make ([]int,10)
for index,_:=range sli{
	sli[index]=rand.Intn(1002)+2
}
for i:=1;i<len(sli);i++{
	for j:=0;j<len(sli)-i;j++{
		if sli[j]>sli[j+1] {
			sli[j], sli[j+1] = sli[j+1], sli[j]
		}
	}
}
return sli
}
