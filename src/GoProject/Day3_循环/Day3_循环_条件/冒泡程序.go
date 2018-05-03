package main
import "fmt"

func main() {
	var a,b,c int
	fmt.Scan(&a)//我已经想到用数组和for循环去命名
	fmt.Scan(&b)
	fmt.Scan(&c)
	if a<b&&a<c{
		fmt.Println(a)
		if b>c{
			fmt.Println(c)
			fmt.Println(b)
		}else{
			fmt.Println(c)
			fmt.Println(b)
		}
	}else if b<c{
		fmt.Println(b)
		if a>c{
			fmt.Println(c)
			fmt.Println(a)
		}else{
			fmt.Println(a)
			fmt.Println(c)
		}
	}else {
		fmt.Println(c)
		if b>a{
			fmt.Println(a)
			fmt.Println(b)
		}else{
			fmt.Println(b)
			fmt.Println(a)
		}
	}
}
