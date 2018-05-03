package main

import "fmt"

func main() {
	var map1 map[int]string
	map1 = map[int]string{1: "wangzhe", 2: "juediq", 3: "lianliankan", 4: "xiaoxiaole"}

	var map2 map[string]string
	map2 = map[string]string{"name": "twodog", "age": "30", "sex": "male", "address": "BeiJing"}
	map2["sex"] = "female"
	fmt.Println(map2)
	fmt.Println(map1)
	for key, _ := range map1 {
		fmt.Println(map1[key])
	}
	for i:=1;i<=len(map1);i++{
		fmt.Print(map1[i])
	}
	fmt.Print("==========================\n")
	fmt.Println(map1==nil,map2==nil)
	fmt.Printf("%p,%p\n",map1,map2)
	var map3 map[int]string
	fmt.Printf("%p",map3) //no address
	//delete(map1,"[1]")
		//mao4:=make(map[string]string)

		var map11 map[int]int
		fmt.Printf("%p\n",map11)
	 	map11=map[int]int{111:11,1211:11}
	 	fmt.Print(map11)
	fmt.Printf("%p\n",map11)


}
