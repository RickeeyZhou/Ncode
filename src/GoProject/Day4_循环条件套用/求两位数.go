package main

import "fmt"

func main() {
	var a int
	var B int
	var b int //b=3+a*10
	for b = 999; b >= 100; b-- {
		for B=0;B<=999;B++ {
			if B-b == 468||b-B==468 {  							//所有相差468的情况
				for a = 0; a <= 99; a++ {
					if b == 3+a*10 &&B==a+300 {
						fmt.Print(a, "\n")
						fmt.Println(B, b)

					}
				}
			}
		}
		//for b=999;b>=0;b-- {
		//	for B = 0; B <= 999; B++ {
		//		if b-B == 468 {
		//			for a = 0; a <= 99; a++ {
		//				if b == 3+a*10 || B == a+300 {
		//					fmt.Print(a, "\n")
		//					fmt.Println(B, b)
		//				}
		//			}
		//		}
		//	}
		//}
		/*循环 10-99

 */
		//for a:=10;a<=99;a++{
		//	if a+300+468==10*a+3||a+300==10*a+3+468{
		//		fmt.Println(a)
		//		fmt.Println(a+300,10*a+3)

	}
}
