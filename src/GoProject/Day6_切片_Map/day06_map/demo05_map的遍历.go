package main

import (
	"fmt"
	"sort"
)

func main()  {
	map2:= map[int]string{1:"王者农药",2:"绝地求生",3:"连连看",4:"传奇霸业",5:"消消乐"}
	fmt.Println(" map的长度：",len(map2))
	//1.定义一个slice
	s1 := make([]int,0,len(map2))
	//2.遍历map获取key-->s1中
	for key := range map2{
		s1 = append(s1, key)
	}
	//3.给s1进行排序
	sort.Ints(s1)//使用sort包下的方法直接排序，不用自己写冒泡了。
	//4. 遍历s1，map
	for _,k:=range s1{ // 先下标，再数值
		fmt.Println(k, map2[k])
	}

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)


}
