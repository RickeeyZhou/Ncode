package main

import "fmt"

func main() {

	fmt.Println("输入用户名和密码：" )
	var name int
	var Password int
	fmt.Scanf("%d%d",&name,&Password)
	fmt.Printf("用户名:%d\n",name)
	fmt.Printf("密码:%d\n",Password)
	fmt.Scanln(&name)
	fmt.Println("name",name)

}
