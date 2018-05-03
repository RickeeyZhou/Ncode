package main
import"fmt"
func main() {
	/*二元一次方程求解
	循环150次
	x=150-y
	在用一个if条件判断
	 */
	 for x:=150;x>=0;x--{
	 	if y:=150-x;y*3==x-10{
	 		fmt.Print(x,y,"\n")
		}
	 }
	 /*四位数是三位数的三倍，[1000,9999]
	 循环8999次，
	 循环条件后三位的值是 原数据的1/3
	  */
	  for c:=9999;c>=1000;c--{
	  	b:=c%1000
	  	if b*3==c{
	  		fmt.Print(c,"\n")
		}
	  }
	  // 循环500次，找出个百十位，
	  //依次判断是不是四，是4 就break
	  //默认符合条件，
	  for zr:=500;zr>=1;zr--{
	  	var a,b,c int
	  	a=zr/100
	  	b=zr%100/10
	  	c=zr%100%10
	  	if a!=4&&b!=4&&c!=4{
	  		fmt.Print(zr,"\n")

		}

	  }
}
