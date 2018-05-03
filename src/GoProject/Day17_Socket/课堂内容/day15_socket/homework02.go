package main

import (
	"fmt"
	"time"
	"sync"
)

func main()  {
	/*
	channel，同步：
		每次只能有一个goroutine操作channel。。
			<- data, <- ch1


	数据的数值：j的值
	 */
	// N-N的生产者和消费者
	ch1 := make(chan int)
	done:=make(chan bool)
	var mutex sync.Mutex
	j := 0
	//生产者的goroutine
	for i:=0;i<5;i++{
		i:=i
		go func() {
			for {
				// g1,g2,g3,g4,g5
				mutex.Lock()
				ch1 <- j//g1, 0
				fmt.Println("生产者：",i,"产生数据：",j)
				time.Sleep(1*time.Second)
				j++ // 1 ,2 ,3, 4
				if j >=  10{
					fmt.Println("生产者：",i,"结束")
					mutex.Unlock()
					break
				}
				mutex.Unlock()
			}
			done <- true
		}()
	}

	//消费者的goroutine
	for i:=0;i<5;i++{
		i := i
		go func() {
			// g6,g7,g8,g9.g10
			for v:=range ch1{
				fmt.Println("\t消费者，",i,"获取数据：",v)
			}
			done <- true
		}()
	}


	//main程序：
	for i:=0;i<5;i++{
		<-done
	}
	close(ch1)

	for i:=0;i<5;i++{
		<-done
	}

	fmt.Println("main..over.....")
}
