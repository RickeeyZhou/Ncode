package main

import (
	_"github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)
type emp struct {
	ename string
	sal  float64
	deptno int
}
// 查询 emp 表中的员工信息
// 员工的名字,工资 ,部门编号
func main(){
	db,_:=sql.Open("mysql" ,"root:2222@tcp(127.0.0.1:3306)/zyj?charset=utf8")
	defer db.Close()
	sli:=make([]emp,0)
	//
	rows,err:=db.Query("select ename ,sal ,deptno from emp where =20 ")
	if err!=nil{
		fmt.Println("查询失败",err.Error())
		return
	}
	//
	for rows.Next(){
		var ename string
		var sal float64
		var deptno int
		if err:=rows.Scan(&ename,&sal,&deptno);err!=nil{
			fmt.Println("获取失败")
		}
		emp1:=emp{ename,sal,deptno}
		sli=append(sli,emp1)

	}
	rows.Close()
	for _,v :=range sli{
		fmt.Println(v)
	}
}