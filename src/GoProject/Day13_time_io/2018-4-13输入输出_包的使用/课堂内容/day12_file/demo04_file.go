package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main()  {
	//1.创建一个目录,仅创建一层
	//err:=os.Mkdir("C:\\Ruby\\pro\\aa", 0666)
	//fmt.Println(err)
	//2.创建一个目录，可以多层
	//err = os.MkdirAll("C:\\Ruby\\pro\\bb\\cc" , 0666)
	//fmt.Println(err)

	//3.创建文件
	//file, err:=os.Create("C:\\Ruby\\pro\\bb\\cc\\ee.txt")
	//fmt.Println(err) //open C:\Ruby\pro\bb\cc\dd\ee.txt: The system cannot find the path specified.
	//fmt.Println(file) //<nil>

	// 4.删除：慎用，慎用，再慎用。。
	// remove C:\Ruby\pro\bb: The directory is not empty.
	//err := os.Remove("C:\\Ruby\\pro\\bb\\cc\\ee.txt") // bb\\cc\\ee.txt
	//fmt.Println(err)
	//删除文件或目录，无论目录中是否有内容
	//err:=os.RemoveAll("C:\\Ruby\\pro\\bb")
	//fmt.Println(err)

	//5.打开文件,当前程序和本地aa.txt文件之间建立了一个链接。。
	file,err:=os.Open("C:\\Ruby\\pro\\aa.txt")
	fmt.Println(err)
	fmt.Println(os.IsExist(err)) //false
	fmt.Println(os.IsNotExist(err)) //true
	fmt.Println(file)
	// 6.关闭文件
	defer file.Close()//断开程序和文件之间的链接，释放资源。
	// 读操作，写数据。。。

	/*
	文件，目录的路径：
	绝对路径，和相对路径
	absoute path
	relative
	 */

	 fmt.Println(filepath.IsAbs("aa.txt"))//false
	 fmt.Println(filepath.IsAbs("C:\\Ruby\\pro\\aa.txt")) //true
	 fmt.Println(filepath.Abs("aa.txt")) //C:\Ruby\workspace\workspace1802\src\day12_file\aa.txt <nil>


/*
mac系统文件路径
/User/ruby/documents/aa.txt

mac下，在选定文件上，command+i，查看文件的信息
	位置：
*/
}
