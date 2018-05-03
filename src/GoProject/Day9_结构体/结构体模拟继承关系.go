package main

import "fmt"

type animal struct {
	foot int
	eyes int
}
type cat struct {
	animal
	name_cat  string
}
type dog1 struct {
	animal
	name_dog string
	host host1
}
type host1 struct{
	sex string
}
func main(){
	tom_cat:=cat{}
	tom_cat.foot=4
	tom_cat.eyes=2
	tom_cat.name_cat="咖啡猫"
	jack_dog:=dog1{animal{4,2},"哈巴狗",host1{"male"}}

	fmt.Println(jack_dog.host.sex,tom_cat)
}