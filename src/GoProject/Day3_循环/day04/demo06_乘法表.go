package main

import "fmt"

func main() {
	/*
1*1=1
2*1=2 2*2=4
3*1=3 3*2=6 3*3=9
...
9*1=9 9*2=18 9*3=27.....9*9=81
	i * j =
 */

	for i := 1;i < 10;i++{
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

}
