package main

import "fmt"
import "math/rand"
import "time"

func main() {
	x := 0
	y := 0
	rand.Seed(time.Now().Unix())
	x = rand.Intn(4) + 1
	y = rand.Intn(5) + 1
	fmt.Println(x, y)
}
