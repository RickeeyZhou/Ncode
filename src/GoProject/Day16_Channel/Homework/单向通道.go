package main
func main(){
	ch1:=make(chan int)
	ch2:=make(chan<-int )
	ch3:=make(<-chan int )

	fun1(ch2)
	fun1(ch1)
	fun2(ch3)
	fun2(ch3)
}
func fun1(ch chan <- int ){

}
func fun2(ch <-chan int ){

}