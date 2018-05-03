package main

import "fmt"

type Book struct {
	BookName string  //书名
	price    float64 //价格
}
type Person struct {
	name string // string
	age  int    // int
	book Book   // 模拟聚合关系
}

type Person4 struct {
	name  string //string
	age   int    //int
	books []Book //Book的切片
}

/*
array
slice

s1:=[]Book{b1,b2,b3...}
s3:=[]int{1,2,3,4}
s2:=make([]Book,len,cap)

s1=append(s1,v1,v2....)

 */

func printInfo(p Person) {
	fmt.Printf("姓名：%s,年龄：%d,所看的书名：《 %s 》，书的价格：%.2f\n", p.name, p.age, p.book.BookName, p.book.price)
}
func printInfo2(p Person4) {
	fmt.Printf("姓名：%s，年龄：%d\n", p.name, p.age)
	if len(p.books) == 0 {
		fmt.Println("\t此人不看书。。")
	} else {
		for i, v := range p.books {
			fmt.Printf("\t第%d本书名：%s，价格：%.2f\n", i+1, v.BookName, v.price)
		}
	}
}

func main() {
	/*
	结构体的嵌套：
		一个结构体作为另一个结构体的字段类型

	 */
	b1 := Book{}
	b1.BookName = "金平里没有梅"
	b1.price = 35.8
	fmt.Println(b1)
	p1 := Person{}
	p1.name = "王二狗" //string
	p1.age = 30     //int
	p1.book = b1    //Book
	fmt.Println(p1)
	printInfo(p1)

	b2 := Book{"十万个为啥", 44.8}
	p2 := Person{"李小花", 18, b2}
	printInfo(p2)

	p3 := Person{
		name: "rose",
		age:  30,
		book: Book{
			BookName: "西游记",
			price:    88.9,
		},
	}
	printInfo(p3)

	p4 := Person{}
	p4.name = "jack"
	p4.age = 28
	//p4.book.BookName="红楼梦"
	p4.book = Book{
		BookName: "红楼梦",
		price:    58.6,
	}
	printInfo(p4)

	fmt.Println("----------------")
	p5 := Person4{}
	p5.name = "花花"
	p5.age = 18
	b3 := Book{"王二狗流浪记", 88.6}
	b4 := Book{"呼啸山庄", 28.6}
	books1 := []Book{b3, b4, Book{"茶花女", 66.4}}
	p5.books = books1
	fmt.Println("姓名：", p5.name, "年龄：", p5.age)
	for i, v := range p5.books { //i：下标，v是Book对象
		fmt.Printf("第%d本书，书名：%s，价格：%.2f\n", i+1, v.BookName, v.price)
	}

	p6 := Person4{"如花", 17, make([]Book, 0, 10)}

	p6.books = append(p6.books, Book{"火影忍者", 88.7})
	p6.books = append(p6.books, Book{"小猪佩奇", 48.7})
	p6.books = append(p6.books, Book{"熊出没", 38.7})
	p6.books = append(p6.books, Book{"天线宝宝", 58.7})
	printInfo2(p6)

	p7 := Person4{}
	printInfo2(p7)




}
