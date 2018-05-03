package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"

)

func main() {
	db, err := sql.Open("mysql", "root:2222@tcp(127.0.0.1:3306)/zyj?charset=utf8")
	if err != nil {
		if err != nil {
			fmt.Println("connetected failure")
			return
		}

		stmt, err := db.Prepare("INSERT INTO emp (empno,ename,job,hiredate,sal)values (?,?,?,?,?)")
		if err != nil {
			fmt.Println("操作失败")
		}

		result, err := stmt.Exec(9529, "hahah", "sb", "2018-02-03", 30.2)
		if err != nil {
			fmt.Println("插入数据失败")
		}

		LastIndexID, err := result.LastInsertId()
		rowsAffected, err := result.RowsAffected()
		fmt.Println("lastInserted", LastIndexID)
		fmt.Println("line", rowsAffected)

		stmt.Close()
		db.Close()
	}
}