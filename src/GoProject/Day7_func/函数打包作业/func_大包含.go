package main
import "fmt"
// var 单个变量 func 函数 type 声明类型
type testInt func(int)bool //声明一个函数类型
func isOdd(integer int)bool{
	if integer%2==0{
		return false
	}
	return true
}
func isEven( inter int )bool{
	if integer%2==0{
		return true
	}
	return false
}
func filter( slice []int,f testInt)[]int{
	var result []int
	for _,value:=range slice{
		if f(value){
			result=append( result,value )

		}
	}
	return result
}
func main(){
	slice:= []int{1,2,3,4,5,7}
	fmt.Println("slice=",slice)
	odd :=filter(slice,isOdd)
	fmt.Println("Odd elsemtn if slice are ",odd)
	even:=filter(slice,isEven)
	fmt.Println("even elements of slice are ",even)

}