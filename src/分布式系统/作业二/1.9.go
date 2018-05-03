package main

import "fmt"

type students struct {
	name string

}
func main(){
	m:= map[string]students{"people":{"jilao"}}
	//m["people"].name=""
	fmt.Println(m["people"])
	fmt.Printf("%T\n",m["people"])
	a:=m["people"]
	fmt.Println(a.name)
	(m["people"]).name="gay"
	m["people"].name="gay"
}