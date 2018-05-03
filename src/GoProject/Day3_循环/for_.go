package main

import (
	"fmt"
)

func main() {
	//i:=0
	//j:=0
	for i := 0; i < 5; i++ {
		j := 0
		for ; j < 5; j++ { //内层循环体 按格式书写 可以减少错误出现的情况
			fmt.Print("1")
		}
		fmt.Println()
	}
	for a := 1; a <= 5; a++ {
		for b := a; b > 0; b-- {
			fmt.Print(`*`)
		}

		fmt.Println()
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {     //乘法表要求 J从1开始，表达式一 固定j:=1
			fmt.Print(i, `*`, j, ` `)

		}
		fmt.Println()
	}
	/*练习5：百元百鸡，一百元钱买100只鸡，公鸡5元一只，母鸡3元一只，小鸡1元3个。*/
	c := 0
	for i := 0; i <= 20; i++ {
		for j := 0; j <= 33; j++ {
			for k1 := 0; k1 <= 99; k1 += 3 {
				if i+j+k1 == 100 && 5*i+3*j+k1/3 == 100 {
					c++
					fmt.Print(i, j, k1, "\n")

				}

			}
		}
	}
	fmt.Print(c)    //百元百鸡

	//素数
	for i := 2; i <= 100; i++ {
		  flag:=true
		for j := i - 1; j > 1; j-- {
			if i%j == 0 {
				flag=false
				break
			}

			//if j == 2 && i%j != 0 {     //相当于标志位  flag
			//	flag++
			//	fmt.Print(i, " is a suhu\n")
			//
			//}
		}
		if flag{
			fmt.Printf("%d,susu\n",i)
		}
	}
	//fmt.Print("\n",flag)
	//百鸡 公鸡5元一只，母鸡3元一只，小鸡1元3个 不超过有100
	flag3:=0
	for sma:=100;sma>=0;sma--{
		for mu:=30;mu>=0;mu--{
			for gon:=20;gon>=0;gon--{
				if 5*gon+3*mu+sma*1/3<100&&sma+mu+gon==100{
					flag3++
					fmt.Printf("sma:%d,mu:%d,gon:%d\n",sma,mu,gon)
				}
			}
		}
	}
	fmt.Print(flag3)

}
