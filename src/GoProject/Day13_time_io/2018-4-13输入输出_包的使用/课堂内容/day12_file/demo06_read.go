package main

import (
	"os"
	"fmt"
)

func main()  {
	/*
	Read([]bs),默认从头开始读取
		光标在最后，再次读取，返回error为io.EOF
	ReadAt([] bs, offset int),从指定位置开始读取
	 */
	 file,_:=os.Open("C:\\Ruby\\pro\\aa.txt")
	 bs := make([] byte ,4)
	 /*
	 1.ReadAt(bs, 0)-->4
	 2.ReadAt(bs,4)-->4
	 3.ReadAt(bs,8)-->2	EOF
	  */
	 count, err:=file.ReadAt(bs,2)
	 fmt.Println(err) //<nil>
	 fmt.Println(count) //4
	 fmt.Println(string(bs[0:count]))

	 count, err = file.ReadAt(bs,7)
	 fmt.Println(err) //EOF
	 fmt.Println(count) //3
	 fmt.Println(string(bs[:count]))
}
