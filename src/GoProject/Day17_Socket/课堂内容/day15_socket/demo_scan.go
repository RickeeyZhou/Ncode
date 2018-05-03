package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	//s1:=""
	//s2:=""
	//fmt.Scan(&s1,&s2)

	//fmt.Scanln(&s1)
	//fmt.Scanln(&s2)
	//fmt.Println(s1)
	//fmt.Println(s2)

	reader1:=bufio.NewReader(os.Stdin)
	//s3,_:=reader1.ReadString('a')
	//fmt.Println(s3)
	bs,flag,_:=reader1.ReadLine()
	fmt.Println(string(bs))
	fmt.Println(flag)

}
