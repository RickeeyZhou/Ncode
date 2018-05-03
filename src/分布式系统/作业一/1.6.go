package main

import (
	"fmt"

)

func calc(index string,a, b int )int {
	ret:=a+b
	fmt.Println(index,a,b,ret)
	return ret
}
func main(){
	a:=1
	b:=2
	defer calc("1",a,calc("10",a,b))
	a=0
	defer calc("2",a,calc("20",a,b))
	b=1
}

/*
line 16:
猜测
20,0,2,2
2,0,2,2

10,1,2,3
1,1,3,4
看结果发生这样的事 是因为先执行了参数的调用,
defer 延迟的参数随函数部分被转到栈内存中
 */
