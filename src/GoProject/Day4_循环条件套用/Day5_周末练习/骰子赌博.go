package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	/*
	开始一局 产生随机数 求和
	判断 该结果 十大还是小
	将结果存入bool 变量 大= true 小= flasfse
	从键盘扫描 玩家 的结果
	默认 玩家赢了 ，一旦输入 小 small 则判断输
	计价系统 总金额  100
	输入 1到20  赢了 +  输了-

	 */
	var duzhu int
	var money= 100

flag:	if money>0 {                       //你是不是穷鬼
		fmt.Println("赌局开始")
		fmt.Println("当前钱包余额：", money)
		fmt.Println("请下注")

	reget:	fmt.Scanln(&duzhu)
		//if duzhu>20||duzhu>money{
		//	if duzhu>20{
		//
		//	}
		//	if duzhu>money{
		//		fmt.Println("余额不足")
		//	}
		//	goto reget  //重新下注
		//}
		switch {
		case duzhu>20:
			fmt.Println("土豪金额太大，庄家玩不起")
			goto reget  //重新下注
		case duzhu>money:
			fmt.Println("余额不足")
			goto reget  //重新下注
		case duzhu>0&&duzhu<=20:
			fmt.Println("下注成功")
		default:
			fmt.Println("我不识字，请输入数字")
			goto reget  //重新下注
		}

		// 产生随机数
		rand.Seed(time.Now().Unix())
		res := rand.Intn(7) + 1 //[1 6] !!!!!!!!!会不会重复执行声明
		res += rand.Intn(7) + 1
		res += rand.Intn(7) + 1 //随机数已经产生 ok
		res1 := ""
		switch res {
		case 11, 12, 13, 14, 15, 16, 17, 18:
			res1 = "大" //大
		default:
			res1 = "小"
		}          //后台处理随机数结果 和玩家无关
		// 玩家输入
		fmt.Println("买大买小，买定离手")
		gamer := ""
		fmt.Scanln(&gamer)
		switch gamer {
		case res1:
			fmt.Printf("恭喜你！赢得%d的赌注\n", duzhu)
			money+=duzhu
			fmt.Printf("你现在的筹码：%d", duzhu+money)
		default:
			fmt.Printf("输掉%d的筹码\n", duzhu)
			money-=duzhu
			fmt.Printf("你现在的筹码：%d", money-duzhu)
		}
		// 是否继续游戏
		fmt.Println("你是否还想继续游戏")
		fmt.Println("	  Yes     No      ")
		var goon string
		fmt.Scanln(&goon)
		switch goon {
		case "yes", "Yes", "YES", "YEs", "YeS":
			goto flag
		case "NO", "No", "no":
			fmt.Println("Game over")

		}
	}else{
		fmt.Println("余额不足")
	}
}
//破产清算