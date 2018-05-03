package main
/* 操场[100,200]人，
三人一组多一人，四人一组多两人，
三人一组，五人一组多三人
思路：循环100次，符合条件的情况输出
*/
import "fmt"

func main() {
	for a:=100;a<=200;a++{
		if a%3==1&&a%4==2&&a%5==3{
			fmt.Println(a)
		}
	}
}
