package main

import "fmt"

func main()  {
	/*
	循环中的控制语句：break，continue
	循环正常结束：循环条件不满足，
	通过break和continue来结束循环。
	break：用于强制结束循环，无论循环条件是否满足
	continue：结束了某一次循环，循环下次继续。
	 */
	 for i := 1;i<=10;i++{
	 	if i % 2==0{
	 		fmt.Println(i)
	 		break
		}
	 }

	 for i := 1;i<=10;i++{
	 	if i == 5{
	 		//break
	 		continue
		}
	 	fmt.Println(i)
	 }
	 // 1.练习1，统计1-100内能被3整除不能被5整除的数，仅要前5个。

	 // 2.打印从1开始的前20个奇数。
	 i:=1
	 count := 0
	 for {
	 	if i % 2 != 0{
	 		fmt.Print(i,"\t") // 1
	 		count++ // 1
	 		if count == 20{
	 			break
			}
		}
		i++ // 2

	}
}
