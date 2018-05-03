package main

import (
	"math"
	"fmt"
)

func main() {


}
// changeable
func myfunc(arg ...int ){

}
for _,n := range arg{
	fmt.Printf ("And the number is : %d\n",n)
}
// 值传递
package main
import " fmt "
import "math "
 func main() {

 getSquareRoot:=func(x float64)float64{
 	return math.Sqrt(x)
}
fmt.Println(getSquarRoot(9))
}
}