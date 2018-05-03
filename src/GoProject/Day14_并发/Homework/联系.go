package main

import (
	"time"
	"fmt"
)

func main(){
	go printLetter()
	go printNum()
	time.Sleep(3000*time.Millisecond)
	fmt.Println("main,over....")
}
func printNum(){
	for i:=1;i<=5;i++{
		time.Sleep(150*time.Microsecond)
		fmt.Println(i,"\t")

	}

}
func printLetter(){
	for i:='A';i<='E';i++{
		time.Sleep(400.*time.Microsecond)
		fmt.Printf("%c\t",i)
	}
}