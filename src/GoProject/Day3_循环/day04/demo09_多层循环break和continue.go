package main

import "fmt"

func main()  {
	out:for i:= 1;i<=5;i++{
		for j:=1;j<=5;j++{
			if j == 3{
				break out// 默认结束的是里层循环
				//continue // 默认结束的是里层循环
			}
			fmt.Println("i:",i," , j:",j)
		}
	}
}
/*
i:1,j:1
i:1,j:2
i:1,j:3
i:1,j:4
i:1,j:5
i:2,j:1
i:2,j:2
i:3,j:3
...
i:5,j:5
 */
