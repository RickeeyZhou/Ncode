package main

import "fmt"

func main(){
	printString("heloo")
	defer printString("Er_dog")
	printString("meme da")
	defer printString("world")
}
func printString(s string ){
	fmt.Println(s)
}
