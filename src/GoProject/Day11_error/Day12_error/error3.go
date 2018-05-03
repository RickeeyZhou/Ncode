package main

import "fmt"

type RectError struct {
	msg string //信息描述
	wid float64
	len float64
}
func  (r *RectError)Error()string {
	return fmt.Sprintf("宽度：%.2f,高度：%.2f,错误提示：%s",r.wid,r.len,r.msg)
}
func rectArea (width,length float64)(float64,error){
	errMsg:=""
	if width<=0{
		errMsg="width cant be small than 0"
	}
	if length<=0{
		if errMsg==""{
			errMsg="length cant be small than 0"
		}else {
			errMsg+=",length cant be small than 0 either "
		}
	}
	// juge the error infomation
	if errMsg!=""{
		return 0, &RectError{errMsg,width,length}
	}
	area:=width*length
	return area,nil
}
func main(){
	area,err:=rectArea(-4,-9)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(area)
	}
}