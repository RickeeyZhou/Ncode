package main

import (
	"fmt"
	"net/http"
)
func main(){
	func  (b ...int){
		fmt.Printf("%T",b)
	}()
	const a= "http://www.zhenai.com/zhenghun"
	resp,_:=http.Get(a)
	fmt.Println(resp)

}
