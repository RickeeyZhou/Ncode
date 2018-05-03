package main

import "sync/atomic"

var value int32
 func SetValue (dalta int32){
 	for {
 		v:=value
 		// var v int32
 		if atomic.CompareAndSwapInt32(&value,v,(v+dalta)){
			break
		}
	}
 }
func main(){
	value=2
	SetValue(2)

}