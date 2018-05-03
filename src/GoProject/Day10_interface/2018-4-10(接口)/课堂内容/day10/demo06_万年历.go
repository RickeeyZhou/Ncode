package main

import "fmt"

/*
万年历：
判断一个年份是否闰年()
求每个月的天数()
求总天数()
 */
func leapYear(year int) bool  {
	if year % 4 ==0 && year % 100 != 0 || year % 400 == 0{
		return true
	}
	return false
}

func getDayOfMonth(year, month int) int {
	switch month {
	case 1,3,5,7,8,10,12:
		return 31
	case 4,6,9,11:
		return 30
	case 2:
		if leapYear(year){
			return 29
		}else {
			return 28
		}
	}
	return 0
}

 //用于求1900年1月1日到给定月的上一个月的总天数
 // 1900,1,1---2000,3,31
 // 整年：1900-1999
 // 2000年1月-3月
func getDays(year, month int) int{
	// year,2000,month:4
	sum := 0
	//1.整年：
	for i:=1900;i<year;i++{
		if leapYear(i){
			sum += 366
		}else {
			sum += 365
		}
	}
	//2.当年的月份
	day := 0
	for i:=1 ;i<month;i++{
		day = getDayOfMonth(year, i)
		sum += day
	}
	return sum
}

func printCalender(year, month int)  {
	//1.计算空格的数量
	days := getDays(year, month)
	kong:= (days+1)%7

	fmt.Println("空格量：",kong)
	//2.计算该月的天数
	day := getDayOfMonth(year, month)
	//3.打印
	//fmt.Println("星期一\t星期二\t星期三\t星期四\t星期五\t星期六\t星期日")
	fmt.Println("星期日\t星期一\t星期二\t星期三\t星期四\t星期五\t星期六")
	//空格
	for i:=0;i<kong;i++{
		fmt.Print("\t\t")
	}
	//数字
	for i:=1;i<=day;i++{
		fmt.Print(i,"\t\t")
		//换行
		if(i+kong) % 7 ==0{
			fmt.Println()
		}
	}
}
func main()  {
	year := 2018
	month := 3
	printCalender(year, month)
}

