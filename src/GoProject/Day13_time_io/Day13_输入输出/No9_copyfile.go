package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	src := "/home/rickeey/测试/aa.jpeg"
	des := "/home/rickeey/测试/haha/bb.jpeg"
	total,err:=copy(src,des)
	if err!=nil{

		 fmt.Println(err)
		 return
	}
	fmt.Print(total,"over" )


}
func copy(srcname, desname string) (int, error) {
	des, _ := os.OpenFile("/home/rickeey/测试/haha/bb.jpeg", os.O_RDWR|os.O_CREATE, 0777)
	src, _ := os.Open("/home/rickeey/测试/aa.jpeg")
	bs := make([]byte, 1024)
	defer des.Close()
	defer src.Close()
	var total int

	for {  ////复制部分设计的很巧妙， 不停的复制的，一个无限循环，考虑到退出条件是，发现这个条件必须在循环体内才出现，
			//所以用if 语句写在里面


		count, err := src.Read(bs)
		if err != nil && err != io.EOF {
			return total, err
		} else if err == io.EOF {
			des.Write(bs[:count])
			fmt.Println("the end of pic")
			break
		} else {
			des.Write(bs)
			total+=count
		}
	}
	return total,nil

}
