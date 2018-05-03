package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	/* 打印这是一个游戏 提示请输入 石头剪刀布
	计算机产生随机数
	判断输赢 思路
	一.将玩家所有输赢情况列出来，一一比对
	二.将玩家

	 */
	var bot int
	var gamer string
flag:
	rand.Seed(time.Now().Unix())
	bot = rand.Intn(4) + 1 //【1，3】

	fmt.Println("这是一个猜拳游戏，请输入你所想的：")
	fmt.Scan(&gamer)
	if gamer=="over"{
		goto over
	}
	if bot == 1 {
		switch gamer {
		case "石头":
			fmt.Println("平了")
		case "剪刀":
			fmt.Println("loose - -")
		case "布":
			fmt.Println("winning >_<")

		default:
			fmt.Println("你说啥")

		}
	}
	if bot == 2 {
		switch gamer {
		case "石头":
			fmt.Println("winning >_<")

		case "剪刀":
			fmt.Println("平了")
		case "布":
			fmt.Println("Loose - -")
		default:
			fmt.Println("你说啥")
		}
	}
	if bot == 3 {
		switch gamer {
		case "石头":
			fmt.Println("loose - -")

		case "剪刀":
			fmt.Println("winning >_<")
		case "布":
			fmt.Println("平了")
		default:
			fmt.Println("你说啥")
		}
	}

goto flag
over:
	fmt.Println("Thank u")

}
