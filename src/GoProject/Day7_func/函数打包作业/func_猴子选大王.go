package main

import "fmt"

func main() {
	fmt.Print("plz tell me how many monkey here?")
	var a int
	fmt.Scanln(&a)
	a = monkey(a)
fmt.Print(a)
}
func monkey(a int) int {
	var sli []int=[]int{}
	for i := 0; i < a; i++ {
		sli=append(sli,1)
	}
	var count int
	var sum int
over:
	for i:=0; ; i++{
		for _, value := range sli {
			if value == 1 {
				count++
				if count == 3 {
					value = 0
					count = 0
					sum++
				}
			}
			if sum == a-1 {
				break over
			}
		}
	}
	f:=0
	for  index,value :=range sli {
		if value ==1 {
			f=index
		}
	}
	return f
}
