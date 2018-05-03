package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main(){
	runtime.GOMAXPROCS(7) //// 控制性能
	wg:=sync.WaitGroup{}
	wg.Add(20) /// 优先等待组
 	for i:=0;i<10;i++{
 		if i==5 {
 			//time.Sleep(5*time.Second)
		}
 		go func() {                      // goroutine 1
 			fmt.Println("i",i)
 			wg.Done()
		}()
	}
	for i:=0;i<10;i++{					 // goroutine 2
		go func (i int ){
			fmt.Println("i:",i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
/*
主函数 创建了一个等待优先组,
组内有20个子协程,
先声明,子协程有打断 主函数执行的可能,我将在i=5 的时间点,睡眠一毫秒,让子协程占用资源,达到验证子协程存在打断主函数的可能.(经过验证主函数十分强势,goroutine一般都在等待主函数完成)
输出结果,  0~9至少出现一次,由goroutine2产生,原因:在goroutine2中将i的值在进入子协程第一时间传递给了自己内部的变量i(同名,改为任意名字也可)
         goroutine 1 本质是由于资源的占用情况, 决定输出什么, 根据主线程强于子线程的原则, 最大可能性是输出 9个 10
 */
 /*
 线程:计算机的处理方式,一堆有序 的而且一个(逻辑)cpu可是彼得命令
 进程:一个正在执行的程序
 一个cpu 同一时刻只能执行一个cpu命令
 物理cpu 逻辑cpu
 一个物理cpu 分割成多个逻辑cpu

 上下文切换: 将命令未执行完毕的线程保存到寄存器.

  */
