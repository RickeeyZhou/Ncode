package main

import "fmt"

func main(){
a:= [10]int {0:1,1:2,2:3,3:4}
b:=[10]float64{0.11,0.222}
//c:=[10]string{"11","2222","222"}
arr1:=[10]int{5,4,3,7,10,2,9,8,6,1}
var sum float64
for i:=1;i<=10;i++{
sum=float64(a[i])+b[i]+float64(arr1[i])
}
fmt.Print(sum)
}