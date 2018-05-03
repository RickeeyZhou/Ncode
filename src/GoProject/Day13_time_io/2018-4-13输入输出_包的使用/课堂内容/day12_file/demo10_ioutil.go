package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func main()  {
	fileName:="c:\\Ruby\\pro\\aa.txt"
	bs, err:=ioutil.ReadFile(fileName)
	fmt.Println(err)//<nil>
	fmt.Println(bs)
	fmt.Println(string(bs))

	fileName2:="C:\\Ruby\\pro\\aoo.txt"
	s1:="HelloWorld"
	err=ioutil.WriteFile(fileName2,[]byte(s1),0777)
	fmt.Println(err)

	s2:="qwertyuiopasdfghjklzxcvbnm"
	reader1:=strings.NewReader(s2)//
	//file,err:=os.Open("")//Read()
	bs2,_:=ioutil.ReadAll(reader1)
	fmt.Println(string(bs2))
}
