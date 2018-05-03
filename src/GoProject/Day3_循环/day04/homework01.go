package main

import "fmt"

func main()  {
/*
模拟登录，键盘上输入用户名和密码，如果用户名是admin密码是123，或者
用户名是zhangsan密码是zhangsan123，都表示可以登录。否则打印登录失败

 */
 username :=""
 password :=""
 fmt.Println("请输入用户名和密码：")
 fmt.Scanln(&username,&password)
 if (username == "admin"&& password =="123") || (username=="zhangsan"&& password =="zhangsan123"){
 	fmt.Println(username,",登录成功")
 }else {
 	fmt.Println("登录失败")
 }

}
