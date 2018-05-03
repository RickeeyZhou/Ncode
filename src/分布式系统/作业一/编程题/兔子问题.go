package main

import "fmt"

// 兔子问题就是递归的问题
func main() {
	var mon int  //时间
	//当月的值等于前两的个值和
	fmt.Scanln(&mon)
	fmt.Println(value(mon))
}
func value (mon int )int {

	if mon==1||mon==2{
		return 1
	}
	return (value(mon-1)+value(mon-2))

}