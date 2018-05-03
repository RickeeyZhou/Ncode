package main

import (
	"fmt"
	"strings"
)
func main() {
	const a="hello hello hello world"
	fmt.Println(strings.Count(a,""))
	fmt.Println(strings.Count(a,"llo"))
	fmt.Println(strings.Index(a,"wor") )

}

