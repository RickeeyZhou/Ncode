package main


func main(){
cir(3)
}
func cir(r1 uint)(float64 error ){
	msg:=""
	if r1<=0{
		msg="半径不能小于0"
	}
	sum:=float64(3.14*r*r)
	if msg==""{
		return sum,nil
	}else{
		return 0,&cirError{msg}
	}
}
type cirError struct{
	inf string
}
func (a *cirError)Error()string{
	return fmt.println(fmt)
}

