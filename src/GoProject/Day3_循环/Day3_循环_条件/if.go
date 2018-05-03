package main

import "fmt"

func main()  {
	fmt.Println("Plz import ur grade")
	Gr:=100
	fmt.Scan(&Gr)
	if Gr>=90{
		fmt.Println("niubi")
	}else if Gr>=80{
		fmt.Println("laji")
	}else if Gr>=70{
		fmt.Println("shy")
	}else{
		fmt.Print("fuck u mom")
	}

}
