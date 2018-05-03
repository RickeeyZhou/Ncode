package main

import "fmt"
import "math/rand"
import "time"

func main() {
	/* random Game
	1.产生一个随机数100~150
	2.输出提示玩家键盘输入一个数
	3.键盘扫描一个数字

	4.验证答案是否正确
		错误提示：回答错误，偏大或偏小，回到步骤2
		正确提示：输出恭喜你，游戏结束
	 */
	rand.Seed(time.Now().Unix())
	num1 := rand.Intn(51) + 100 //[100,150]
	flag:fmt.Println("随机数已经产生，请您输入一个[100,150]的整数")
	var res int
	fmt.Scanln(&res)
	if res==num1{
		fmt.Printf("恭喜你！回答正确")
	}else {
		switch {
		case num1>res:
			fmt.Println("偏小")
		case num1<res:
			fmt.Println("偏大")
		}
		goto flag
	}
}
