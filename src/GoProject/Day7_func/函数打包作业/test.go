package main

import "fmt"

func main() {
	//ar()
	no4(23)
}

// func ar() {
// 	var arr1 [100]int
//
// 	for index, value := range arr1 {
// 		value = index + 100
//
// 		if value%3 == 1 && value%4 == 2 && value%5 == 3 {
// 			fmt.Println(value)
// 		}
// 	}
// }
func no4(a int) {

	for i := a; i <= 500; i++ {
		if i < 100 && (i < 40 || i > 49) {
			if i%10 != 4 {
				fmt.Print(i, "\t ")
			}
		}
		//		} else if i < 400 { //a>=100 a<400
		//			if i%100/10 != 4 || i%100%10 != 4 {
		//				fmt.Print(i, "\t ")
		//			}
		//		}
		//	}
		//}
	}
}