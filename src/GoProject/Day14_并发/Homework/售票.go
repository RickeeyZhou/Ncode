package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var ticket = 10
var wg1 sync.WaitGroup

func main() {

	wg1.Add(1)
	go saleUnit("A")
	go saleUnit("B")
	go saleUnit("C")
	go saleUnit("D")
	wg1.Wait()

}
func saleUnit(name string) {
	rand.Seed(time.Now().UnixNano())
	for ; ticket > 0; {
		ticket--
		fmt.Printf("%s下单成功，余量：%d\n", name, ticket)
		if ticket == 0 {
			wg1.Done()
			fmt.Println("No more ticket")
			break
		}
		time.Sleep(time.Duration(rand.Intn(1000) + 1))
	}

}
