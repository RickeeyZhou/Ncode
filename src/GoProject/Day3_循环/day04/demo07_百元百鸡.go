package main

import "fmt"

func main()  {
	/*
	练习5：百元百鸡，一百元钱买100只鸡，公鸡5元一只，母鸡3元一只，小鸡1元3个。
	公鸡：0-20
	母鸡：0-33
	小鸡：100-公鸡-母鸡

	 */
	 for i := 0 ;i<= 20;i++{//公鸡
	 	for j:=0;j<=33;j++{//母鸡
	 		k := 100-i-j // 小鸡的数量
	 		if i * 5 + j * 3 + k / 3 == 100 && k % 3 == 0{
	 			fmt.Println("公鸡:",i,"母鸡:",j,"小鸡:",k)
			}

		}
	 }

}
