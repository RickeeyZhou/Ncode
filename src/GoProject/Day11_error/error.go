package main

import "fmt"

func main(){
	var a,b float64
fmt.Scanln(&a,&b)
fmt.Println(div(a,b))
}


func div(a ,b float64 )float64{
	var c float64
err:=checkmother(b)

if err!=nil {
	fmt.Println(err)
	return 0.0
}
	c=a/b
	return c
}


func checkmother(b float64)error{
	if b==0{
		err:=fmt.Errorf("分母不能为0")
		return err

	}
	fmt.Println("分母合法")
	return nil

}
