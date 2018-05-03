package main

import "fmt"

func main()  {
	const (
		a = iota  // 0
		b          //1
		c          //2
		d = "ha"   // ha  iota =3
		e          // ha
		f = 100    // 100
		g          //100
		h = iota   //7
		i          //8
	)
	const(
		j = iota  // 0
	)
	fmt.Println(a,b,c,d,e,f,g,h,i,j)

	const (
		x = 'A' //65  iota=0
		y // 65 iota=1
		m = iota // 2
		n // 3
	)
	fmt.Println(x, y, m, n)
}
