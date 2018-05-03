package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var ticket11 = 1000
var wg11 sync.WaitGroup
var LL sync.Mutex
func main() {

	wg11.Add(1)
	go saleUnit11("A")
	go saleUnit11("B")
	go saleUnit11("C")
	go saleUnit11("D")
	wg11.Wait()

}
func saleUnit11(name string) {
	rand.Seed(time.Now().UnixNano())
	for  ; ; {
		if ticket11>0 {
			LL.Lock()
			if ticket11 == 0 {
				wg11.Done()
				LL.Unlock()
				fmt.Println("No more ticket")
				break
			}
			ticket11--
			fmt.Printf("%s下单成功，余量：%d\n", name, ticket11)
			if ticket11 == 0 {
				wg11.Done()
				LL.Unlock()
				fmt.Println("No more ticket")
				break
			}
			LL.Unlock()
			time.Sleep(time.Duration(rand.Intn(1000) + 1))
		}

	}

}
