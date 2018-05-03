package main
type Param map[string]interface {}
type Show struct {
	Param //切片作为属性
}
func main(){
	s:=new(Show)

	s.Param=make(map[string]interface{})
	s.Param["rmb"]=10000
}
