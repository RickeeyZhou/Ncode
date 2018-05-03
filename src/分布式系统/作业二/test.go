package main

import "fmt"

type people12 struct{
	name string
}
func ( p *people12) String()  {
	 fmt.Printf("%T",fmt.Sprintf("print:%v",p))

}
func main(){
	p:=&people12{"sb"}
	p.String()
}
// p