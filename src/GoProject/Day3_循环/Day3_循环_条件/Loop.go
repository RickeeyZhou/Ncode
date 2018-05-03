package main

import (
	"fmt"
)

func main() {
	var day,temp int
	fmt.Println("Please import a num")
	fmt.Scanln(&day,&temp)
	switch day{
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("turseday")
	case 3:
		fmt.Println("wednseday")
	case 4:
		fmt.Println("thurseday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturdaut")
	}
	if temp==1{
		fmt.Println("Mon")
	}else if temp==2{
		fmt.Println("tues")
	}else if temp==3{
		fmt.Println("wedn")
	}else{
		fmt.Println("holiday")

	}

}

