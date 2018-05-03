package main

import "fmt"

func main()  {
	/*
	统计1月-5月的总天数。
	 */
	 year := 2018// 年份
	 sum := 0 // 总天数
	 day := 0 // 求每个月天数
	 n := 8
	 for i:= 1;i<=n;i++{
		 switch i {
		 case 1, 3, 5, 7,8,10,12:
		 	day = 31
		 case 4, 6,9,11:
		 	day = 30
		 case 2:
		 	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0{
		 		day = 29
			}else {
				day = 28
			}
		 }
	 	sum = sum + day

	 }
	 fmt.Printf("1-%d月的天数：%d",n,sum )

}
