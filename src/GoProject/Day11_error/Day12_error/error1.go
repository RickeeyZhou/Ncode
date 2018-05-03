package main

import (
	"fmt"
	"errors"
)

func main(){
	var err error
	fmt.Println(err)
	err=errors.New("test")
	fmt.Println(err)
	fmt.Printf("%T\n",err)

	err2:=checkAge(-30)
	fmt.Println(err2)
	fmt.Println("%T\n",err2)
	if err2!=nil{
		fmt.Println(err2)
		return
	}
}
func checkAge(age int )error{
	if age <0{
		err:=fmt.Errorf("my old is %d,",age)
		return err
	}
	fmt.Println("old is liag",age)
	return nil
}