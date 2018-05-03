package main

import (
	"fmt"
	"errors"
)

func main()  {
	/*
	error：Go中表示不正常的错误的类型。
		error，是interface类型，表示错误的
			errors.New(string)-->error
			fmt.Errorf(format)-->error
	 */
	 var err error
	 fmt.Println(err) //<nil>
	 err = errors.New("写着玩的")
	 fmt.Println(err)
	 fmt.Printf("%T\n",err) //*errors.errorString
	 // 定义一个函数，用于检验年龄，如果为负数，返回一个error

	 err2 := checkAge(-30)
	 fmt.Println(err2)
	 fmt.Printf("%T\n",err2)
	 if err2 == nil{//无错误
	 	//...后序的操作
	 }
	 if err2 != nil{
	 	fmt.Println(err2)
	 	return
	 }
}
//定义函数，检验年龄，
func checkAge(age int) error  {
	if age < 0{
		//返回error
		//return errors.New("年龄不能为负数")
		err := fmt.Errorf("您的年龄是：%d，不能为负数，所以不合法",age)
		return err
	}
	fmt.Println("年龄是合法的，", age)
	return nil
}
/*
键盘输入两个数：考虑除0情况
定义一个函数，用于求两个数的商，和余数，
 */
