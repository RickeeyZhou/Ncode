package main
import "fmt"

func main() {
	year:=2018
	fmt.Scanln(&year)
	switch year%4{
	case 0:
		fmt.Println("runnian")
	default:
		fmt.Println(`bushi`)
	}
	if year1:=2018;year1%4==0&&year1%100!=0 {
		fmt.Println("runnian")

	}else if year1%400==0{
		fmt.Println("runnian")
	}else{
		fmt.Println("nothing")
	}
}
