package main

import (
	"fmt"
	"time"
)

func main() {
	defer_call()
}
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() {
		time.Sleep(5*time.Second)
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

// defer 的数据结构是栈,性质是先进后出
/*在此次题的运行过顺序是
						子线程 : panic ("异常"),终止主程序,转向栈内存中,栈中存储着 defer 语句,但是异常打印位置无法确定,我猜测是因为"异常打印"和 defer打印也不在一个进程中
defer ......(打印后)
defer ......(打印中)
defer ......(打印前)
 */
/* 输出: 触发异常
		打印后
		 打印中
		 打印前
异常可以在打印语句的任意位置
*/
