package main

import "fmt"

func main()  {
	fmt.Printf("啦啦啦%d,%.2f,%s",1,3.1415,"么么哒")
	s1 :="啦啦啦"+"1"+"3.1415"+"么么哒" //strconv
	fmt.Print(s1)
	//error
	// errors.New(string)
	err:=fmt.Errorf("啦啦啦%d,%.2f,%s",1,3.1415,"么么哒")
	fmt.Printf("%T\n",err)
}
