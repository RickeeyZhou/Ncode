package main

import "fmt"

func main(){
sli:="asdfg"
arr:=[3]int{2,3,2}
sli2:=sli[2:] ///依旧是数组
sli3:=arr[:] ///是切片
fmt.Println(sli2)
fmt.Printf("%T\n,,,,,,%T\n",sli2,sli3)
//backout(sli2)
}
func backout( sli string){
	 	 for _,value:=range sli{
	 	defer 	fmt.Printf("%c ",value)
		}

}