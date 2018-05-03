package main

import "fmt"

func main() {
	n := 7
	count := 0
	var arr [7]int = [7]int{1, 1, 1, 1, 1, 1, 1}
	for j := 1; j <= n-1; j++ {
		for i := 0; i < n; i++ {
			if arr[i] == 1 {
				count += 1
				if count%3 == 0 {
					arr[i] = 0

				}
			}

		}
		if j == 6 {
			fmt.Print(arr, count, "\n")
		}
		if j == 7 {
			fmt.Print(arr, count, "\n")
		}
	}

	fmt.Print("flag")
	fmt.Print(arr)
	for j := 0; j < n; j++ {
		if arr[j] == 1 {
			fmt.Println(j)
		}
	}
}
