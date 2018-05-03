package 接口

import "fmt"

type IEngine interface {
	start()
	speedup()
	stop()
}
type Yamaha struct {
	Name string
}
type Honda struct{
	Name string
}
func (s Yamaha)start(){
	fmt.Println("Start:",60)
}
func (s Yamaha)speedup(){
	fmt.Println("Speedup:",80)
}
func (s Yamaha)stop(){
	fmt.Println("Stop:",0)
}
func (s Honda)start(){
	fmt.Println("Start:",40)
}
func (s Honda)speedup(){
	fmt.Println("Speedup:",120)
}
func (s Honda)stop (){
	fmt.Println("Stop:",0)
}
type   Car1 struct{
	IEngine // 接口类型
}
func ( cc Car1) TestIEngine(){
	cc.start()
	cc.stop()
	cc.speedup()
}

	//类型
	//int
//func ( s car1)testIEngine(){
//	s.stop()
//	s.speedup()
//	s.start()
//}
