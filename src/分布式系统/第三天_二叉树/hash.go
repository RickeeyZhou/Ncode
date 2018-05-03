package main

import (
	"crypto/sha256"
	"fmt"
)

func main(){
	const(
		input1="hello"
		input2="heeee"

	)
	first:=sha256.New()
	first.Write([]byte(input1))
	fmt.Printf("%x\n",first.Sum(nil))



}
