package main

import "Homework/Day10_接口/接口"

func main(){
	mycar:=接口.Car1{IEngine:}
	mycar.TestIEngine()
	ss:=接口.Car1{IEngine:接口.Honda{}}

	//Jerome:=接口.Honda{"Nissan"}
	//var TT 接口.IEngine
	//TT=mycar
	//TT.TestIEngine()





}
