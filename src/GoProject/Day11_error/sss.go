package main

import (
	"errors"
	"fmt"
)

func main(){
	error：//Go中表示不正常的错误的类型。
	error，//是interface类型，表示错误的
	errors.New(string)-->error
	fmt.Errorf(format)-->error
}