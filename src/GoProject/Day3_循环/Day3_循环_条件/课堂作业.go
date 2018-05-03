package main

/*练习1：打印1-100内，所有的偶数
	练习2：打印1-100内，能被3整除，不能被5整除的数，每行打印5个
	统计打印数字的个数。
	练习3：求1-100内，奇数的和*/
import "fmt"

func main() {
	for a := 2; a%2 == 0 && a <= 100; a += 2 {
		fmt.Println(a)
	}
	c:=0
	for b := 1; b <= 100; b += 1 {
		if b%3 == 0 && b%5 != 0 {
			fmt.Print(b," ")
			c++
			if c%5==0{
				fmt.Print("\n")
			}

		}
		//fmt.Println("\n",c)
	//	fmt.Println(c)

	}
	fmt.Println("\n",c)
	sum:=0
	for d:=0;d<=100;d++{
		if d%2!=0{
			sum+=d
		}
	}
	fmt.Println(sum)
}
