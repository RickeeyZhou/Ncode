package main

import (
	"fmt"
	"crypto/sha256"
)

func isvaliddifficlt1 (hash string,difficulty int )bool{
	b:=len(hash)
	var i int
	for i=0;i<b;i++{
		if hash[i]!='0'{
			break
		}
	}
	return i>=difficulty
}
func main() {
ss:="00009ab51b8006c4a23361912c1c857f65450ce565829bd0172739368e7eaba4"
t:=isvaliddifficlt1(ss,4)
fmt.Println(t)
fmt.Println(ss)
x:=fmt.Sprintf("%x",ss)
fmt.Println(x)

sha:=sha256.New()
sha.Write([]byte(ss))
a:= sha.Sum(nil)
fmt.Println(a)
b:=fmt.Sprintf("%x",a)
fmt.Println(b)
fmt.Println(sha.Size())

}
