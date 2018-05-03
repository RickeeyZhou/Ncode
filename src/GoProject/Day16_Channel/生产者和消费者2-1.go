package main

import "fmt"

func main(){
	done:=make(chan bool)
	ch1:=make(chan int )

	go func(){
		for i:=1;i<100;i++{
			ch1<-i
			fmt.Println("No.1 正在生产")
		}

	}()
	go func(){
	for i:=100;i<200;i++{
		ch1<-i
		fmt.Println("No.2 正在生产")
	}

	}()
	go func() {
		for{
			 for n:=range ch1{
				fmt.Println("消费者 产品编号：",n)
			}
		}
		done<-true
	}()


	<-done
	<-done

	<-done
}