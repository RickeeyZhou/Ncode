package main

import "fmt"

func main() {
	fmt.Printf("hahah %d,%.2f,%s", 1, 3.1414926, "memeda0")

	s1 := "hhahhah" + "1" + "11222" + "mememda"
	fmt.Println(s1)

	err := fmt.Errorf("lalalal%d,%.2f,%s", 1, 3.1415, "memeda")
	fmt.Printf("%T\n,%v",err,err)
}
