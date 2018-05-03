package main

import "fmt"

func main()  {
	// N-N的生产者和消费者
	ch1 := make(chan int)
	done:=make(chan bool)



	//生产者的goroutine
	for i:=0;i<5;i++{
		i:=i
		go func() {
			for j:=0;j<100;j++{
				ch1 <- j
				fmt.Println("生产者：",i,"产生数据：",j)
			}
			done <- true
		}()
	}

	//消费者的goroutine
	for i:=0;i<5;i++{
		i := i
		go func() {
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
