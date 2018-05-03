package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	var s[10] int
	rand.Seed(time.Now().Unix())
	for i:=0;i<len(s);i++{

		s[i]=rand.Intn(10)
		fmt.Print(s[i],"\t")
	}
	for i:=1;i<len(s);i++{
		for j:=0;j<len(s)-i;j++{
			if s[j]<s[j+1]{
				s[j],s[j+1]=s[j+1],s[j]
			}

		}
	}
	fmt.Print(s)
}
