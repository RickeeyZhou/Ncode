package main

import (
	"sync"
	"fmt"
)

var wait1 sync.WaitGroup
var Sum = 5000

func main() {
	wait1.Add(1)
	var a, b, c, d int

		go a_sale(&a)
		go a_sale(&b)
		go a_sale(&c)
		go a_sale(&d)



	wait1.Wait()
	fmt.Println()
	fmt.Println(a, b, c, d)
	fmt.Println(a+ b+c+ d)

}
func a_sale( a *int) {
	defer wait1.Done()
	for ;Sum>=1;Sum--{
		*a++

	}
	fmt.Println("over")









	//for {
	//	if Sum >= 1 {
	//		Sum --
	//		*a = *a + 1
	//
	//		time.Sleep(time.Nanosecond)
	//		//if Sum == 0 {
	//
	//	} else {
	//		wait1.Done()
	//		return
	//	}
	//}
	//
	//
	//

}

//func b_sale() {
//	wait.Done()
//}
//func c_sale() {
//	wait.Done()
//}
//func d_sale() {
//	wait.Done()
//}
