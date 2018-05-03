package main
import "fmt"
func main() {
	/*万年历
	怎么打印？
	第一行输出星期天到星期一
	循环一个月的天数
	循环遇到星期六就回车*/
	var year, month int
	flag:fmt.Scanf("%d,%d", &year, &month)
	fmt.Printf("年份：%d\n月份：%d\n", year, month)
	//x := 7 //第一天星期几
	//y := 30
	var x, y int
	var Day1 int //当年的天数
	var Day int // 1900到今年1.1的天数 //根据年份和月份得到 x y
	if (year%4 == 0) && (year%100 != 0)||year%400==0 { //是闰年
		//Day := ((year-1900)/4)*(365+365+365+366) + ((year-1900)%4)*365 //1900.1.1总天数
		for k:=1900;k<year;k++{
			if (year%4 == 0 && year%100 != 0)||year%400==0{
				Day+=366
			}else{
				Day+=365
			}
		}
		for k := 1; k <month; k++ {
			switch k {
			case 1, 3, 5, 7, 8, 10, 12:
				Day1 += 31
			case 4, 6, 9, 11:
				Day1 += 30
			case 2:
				Day1 += 29
			}
		}
		Day+=Day1

		//if Day%7==0{
		//	x=0
		//}else{
		//	x=Day%7
		//}
		x=(Day+1)%7
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			y = 31
		case 4, 6, 9, 11:
			y = 30
		case 2:
			y = 29
		}
	} else { //不是闰年
		//Day := ((year-1900)/ 4) * (365 + 365 + 365 + 366) //1900.1.1总天数
		for k:=1900;k<year;k++{
			if (year%4 == 0 && year%100 != 0)||year%400==0{
				Day+=366
			}else{
				Day+=365
			}
		}
		for k := 1; k < month; k++ {
			switch k {
			case 1, 3, 5, 7, 8, 10, 12:
				Day1 += 31
			case 4, 6, 9, 11:
				Day1 += 30
			case 2:
				Day1 += 28
			}
		}
		//Day+=Day1
		//if Day%7==0{
		//	x=0
		//}else{
		//	x=Day%7
		//}
		x=(Day+1)%7 //0,1,2,3,4,5,6
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			y = 31
		case 4, 6, 9, 11:
			y = 30
		case 2:
			y = 28
		}

	}
	fmt.Println(x, y)
	fmt.Println(Day)
	fmt.Print("星期日 ")
	fmt.Print("星期一 ")
	fmt.Print("星期二 ")
	fmt.Print("星期三 ")
	fmt.Print("星期四 ")
	fmt.Print("星期五 ")
	fmt.Println("星期六 ")

	//i := 1
	//	for ; i < x; i++ {
	//		fmt.Print("     ") // 打印空格
	//	}
	//	//i ==x
	//	for day := 1; day <= y; day++ { // y是一个月的天数,day打印计数器
	//		fmt.Print(day)
	//		if i%7 == 0 { //打回车
	//			fmt.Println()
	//		}
	//
	//	}
	//}
	for b := 1; b <x; b++ { //x,0,1,2,3,4
		fmt.Print("       ") //打空格
	} //b == x

	for day := 1; day <= y; day++ {
		if day < 10 {
			fmt.Print(day, "      ")
		}
		if day >= 10 {
			fmt.Print(day, "     ")
		}
		if (x+(day)-1)%7==0 {
			fmt.Println()
		}
	}
goto flag
}
