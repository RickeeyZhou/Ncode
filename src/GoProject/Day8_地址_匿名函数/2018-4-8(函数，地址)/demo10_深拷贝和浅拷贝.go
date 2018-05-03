package main

import "fmt"

func main() {
	/*
	深拷贝和浅拷贝
		深拷贝：复制的是数据
		浅拷贝：赋值的是地址

	数据类型的分类：
		值类型：复制的是数据本身
			基本类型，数组
		引用类型：复制的是地址
			切片，map，函数。。
	 */
	a := 10
	b := a //将a赋值给b，将a的数据复制一份，给b
	fmt.Println(a, b)
	b = 20
	fmt.Println(a, b)
	fmt.Printf("%p,%p\n", &a, &b)

	c := &a
	fmt.Println(a, *c)
	*c = 20
	fmt.Println(a, *c)

	fmt.Println("-------------------")
	// 深拷贝：
	arr1 :=[4]int{1,2,3,4}
	fmt.Println(arr1) //[1 2 3 4]
	arr2 := arr1//将arr1的数值复制了一份，给arr2
	fmt.Println(arr2) //[1 2 3 4]
	arr2[0] = 100
	fmt.Println(arr1) //[1 2 3 4]
	fmt.Println(arr2) //[100 2 3 4]

	p1 := &arr1 //取arr1的地址，p1，指针
	fmt.Printf("%T\n", p1) //*[4]int
	(*p1)[0] = 200
	fmt.Println(arr1) //[200,2,3,4]
	fmt.Println(*p1)

	fmt.Println("------------")
	//切片：长度和容量
	// 切片是引用类型，默认是浅拷贝，复制的是切片的地址
	s1 :=[]int{1,2,3,4} //xc04200c4000
	fmt.Println(s1) // [1 2 3 4]
	s2:=s1 // 将s1的地址给s2
	fmt.Println(s2) // [1 2 3 4]
	s2[0] = 100
	fmt.Println(s1) // [100 2 3 4]
	fmt.Println(s2) // [100 2 3 4]
	fmt.Printf("%p,%p,%p\n",s1,&s1,&s2)

	//深拷贝：切片，遍历s1的每个数据，追加到s3中。或者：copy(dest,src)
	s3 := make([]int,0,4)
	fmt.Println(s3)
	for i:=0;i<len(s1);i++{
		s3 = append(s3, s1[i])
	}
	fmt.Println(s3)
	s3[0] = 1000
	fmt.Println(s1)
	fmt.Println(s3)

	/*
	思考题：
	s1:=append(s1,el)
	 */
}
