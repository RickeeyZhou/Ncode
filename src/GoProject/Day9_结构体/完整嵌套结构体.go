package main

import "fmt"

type Person5 struct{
	name string
	age int
	books []Book1
}
type Book1 struct {
	BookName string
	price float64
}
func printInfo2 (p Person5){
	fmt.Printf("name:%s,age:%d\n",p.name,p.age )
	if len(p.books)==0{
		fmt.Println("\t No Book")
	}else{
		for i,v :=range p.books{
			fmt.Printf("\tNo.%d book:%s,Price:%.2f\n",i+1,v.BookName,v.price)
		}
	}
}
func main(){

	p9:=Person5{}
	p9.name="cnm"
	p9.age=22
	b3:=Book1{"wanerde",2323}
	b4:=Book1{"sdad",232323}
	books2:=[]Book1{b3,b4,Book1{"sss",211}}
	p9.books=books2
	printInfo2(p9)

}
