package main

import (
	"fmt"
	"time"
)

func main()  {
	a := 1
	go func() {
		a = 2
		fmt.Println("子goroutine。。",a)
	}()
	a = 3
	time.Sleep(1)
	fmt.Println("main goroutine。。",a)
}
/*
go run xx.go
C:\Ruby\workspace\workspace1802\src\day13_并发>go run -race demo05_资源的竞争.go

==================
main goroutine。。 3
WARNING: DATA RACE
Write at 0x00c04200e098 by goroutine 6:
  main.main.func1()
      C:/Ruby/workspace/workspace1802/src/day13_骞跺彂/demo05_璧勬簮鐨勭珵浜?go:
11 +0x46

Previous write at 0x00c04200e098 by main goroutine:
  main.main()
      C:/Ruby/workspace/workspace1802/src/day13_骞跺彂/demo05_璧勬簮鐨勭珵浜?go:
14 +0x95

Goroutine 6 (running) created at:
  main.main()
      C:/Ruby/workspace/workspace1802/src/day13_骞跺彂/demo05_璧勬簮鐨勭珵浜?go:
10 +0x84
==================
子goroutine。。 2
Found 1 data race(s)
exit status 66

 */
