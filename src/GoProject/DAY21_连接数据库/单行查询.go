package main

import (
	"database/sql"
	"fmt"
	_"github.com/Go-SQL-Driver/MySQL"
)



func main(){
 db,err:=sql.Open("mysql","root:2222@tcp(127.0.0.1:3306)/zyj?charset=utf8")
 if err!=nil{
 	fmt.Println(err.Error())
 }
 defer db.Close()

 row:=db.QueryRow("select * from emp where ename=?","smith")
 var ename ,job string

 var empno int
 err1:=row.Scan(&empno,&ename ,&job)
 if err1!=nil{
 	fmt.Println("查无数据")
 }else{
 	fmt.Println(empno,ename,job)
 }
db.Close()
}
