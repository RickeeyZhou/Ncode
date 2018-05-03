package main

import "fmt"

func main()  {
	/*
	满足以下两点中任意一点就是闰年

	A：能被4整除，但是不能被100整出。

	B：能被400整除。
	 */
	 year := 2018
	 if year % 4 == 0 && year % 100 != 0 || year % 400 == 0{
	 	fmt.Println(year, ",是闰年")
	 }else {
	 	fmt.Println("是平年。。")
	 }


	 sex := true //true，男，false，女
	 if sex { // sex == true
	 	fmt.Println("是爷们儿就下100层。。")
	 }

	 sex2 := true
	 age := 20
	 if !sex2 && age >= 18{//sex2 == false
	 	fmt.Println("女大十八变。。")
	 }else{
	 	//
	 }
}
