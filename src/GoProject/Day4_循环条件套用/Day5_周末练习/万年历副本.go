package main
import "fmt"
func main() {
	/*万年历
	怎么打印？
	第一行输出星期天到星期一
	循环一个月的天数
	循环遇到星期六就回车*/
	//问题在天数上
	// 接受键盘值的变量
	var year, month int
flag:fmt.Scanf("%d,%d", &year, &month)
	//检查是否接收数据成功
	fmt.Printf("年份：%d\n月份：%d\n", year, month)
	// x：&year年&month月的第一天的星期几设为x
	//y:&month月为y天
	var x, y int
	//Day：1900年到&year 年的天数
	//Day1：&year 当前年份已经过去的天数
	var Day1 int
	var Day int
	//是闰年
	if (year%4 == 0) && (year%100 != 0)||year%400==0 {
		//求Day的天数
		for k:=1900;k<year;k++{
			if (year%4 == 0 && year%100 != 0)||year%400==0{
				Day+=366
			}else{
				Day+=365
			}
		}
		//当前时间距离当前年份一月一日的天数
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
		//此时Day 代表总天数
		Day=Day+Day1
		//总天数+1除7的模为 当月第一天星期几
		if Day%7==0{
			x=1}else{
			x=Day%7+1
		}
		//当前月份的天数
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			y = 31
		case 4, 6, 9, 11:
			y = 30
		case 2:
			y = 29
		}
	} else {             //不是闰年
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
		Day=Day+Day1
		if Day%7==0{
			x=1}else{
				x=Day%7+1
		}
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			y = 31
		case 4, 6, 9, 11:
			y = 30
		case 2:
			y = 28
		}
	}
	//检验是否得到 x y  x为星期几 y为当月天数
	fmt.Println(x, y)
	fmt.Println(Day)
	fmt.Println(Day1)
	//
	fmt.Print("星期日\t")
	fmt.Print("星期一\t")
	fmt.Print("星期二\t")
	fmt.Print("星期三\t")
	fmt.Print("星期四\t")
	fmt.Print("星期五\t")
	fmt.Println("星期六\t")

	//打印第一天之前的空格 假如第一天是星期六，就该打印6个空格
	for b := 1; b <=x; b++ {
		fmt.Print("    \t")
	}
	//打印天数

	for day := 1; day <= y; day++ {


		if day < 10 {
			fmt.Print(day, "     \t")
		}
		if day >= 10 {
			fmt.Print(day, "    \t")
		}
		if (x+(day))%7==0 {
			fmt.Println()
		}

	}
	goto flag
}

