package main

import "fmt"
import "math"

func main()  {
	for i:=10;i<100;i++{
		/*
		30--->
			29,19,7
			m, >  n
			329, 293
			319,193
			307,73
			3xx > xx3
		  m - n == 468

			67,86,59
			m		n
			367 , 	673
			386,	863
			358,	583
			n - m == 468

		85:385,853
		 */
		m := i + 300
		n := i*10+3
		if n - m == 468 || m-n == 468{ // abs(m-n)
		//if math.Abs(float64(m-n)) == 468{
			fmt.Println(i)
		}
	}
}
