package main

import (
	"fmt"

)

func main() {
	const Mon =6
	if Mon<=5&&Mon>=3{
		fmt.Println("spring")
	}else if Mon==6||Mon==7||Mon==8{
		fmt.Println("summer")
	}
	const tem=12
	switch tem{
	case 3,4,5:
		fmt.Println("spring")
	case 6,7,8:
		fmt.Println("summer")
	case 9,10,11:
		fmt.Println("wutu")
	case 12,1,2:
		fmt.Println("winter")
	}
}
