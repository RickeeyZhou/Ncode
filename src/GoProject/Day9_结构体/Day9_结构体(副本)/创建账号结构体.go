package main

import "fmt"

type User struct {
	username string
	password int
}
func main(){
fuzilong:=User{ "Jerome",123456}
zhouyujing:=User{"Rickeey_zhou",1304030304}
fmt.Printf("%T\n",fuzilong.username)
fmt.Println(zhouyujing)
}

