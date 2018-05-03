package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main(){
	err:=os.Mkdir("/home/rickeey/测试/haha",0777)
	fmt.Println(err)
	err=os.MkdirAll("/home/rickeey/测试/SB/SBSBSB",0777)
	/// create a dir /

	os.Create("/home/rickeey/测试/haha/sex.avi")
	file,err:=os.Create("/home/rickeey/测试/haha/sex2.avi")
	fmt.Println(err)
	fmt.Println(file)
	// 关于接口实现对象 ；返回接口类型时，实际返回的是实现类的对象，并且是他的方法，而且
	//如果在方法中是通过地址调用的话，那就是直接返回 对象的地址。
	//返回接口类型时 ，到底是返回对象还是方法。结论是对象，返回地址，...额，发现别的东西还不好使，因为没有名字了，总不能使用类型名字吧

	os.Remove ("/home/rickeey/测试/sex.avi")
	os.RemoveAll("/home/rickeey/测试/SB")

	file2,err:=os.Open("/home/rickeey/测试/ToDoList.sh")
	fmt.Println(err)
	fmt.Println(file2)
	defer file2.Close()
	fmt.Println(os.IsExist(err))
	fmt.Println(os.IsNotExist(err))

	fmt.Println(filepath.IsAbs("sex2.avi"))
	fmt.Println(filepath.IsAbs("/home/rickeey/测试"))
	fmt.Println(filepath.Abs("sex2.avi"))
}