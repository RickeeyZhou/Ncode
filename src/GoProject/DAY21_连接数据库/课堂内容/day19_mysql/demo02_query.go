package main
//step1：导入包
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
type Emp struct {
	Empno int
	Ename string
	Job string
	Hiredate string
	Sal float64
	Deptno int
}

func main()  {
	/*
	查询操作：
	 */
	 //step2:打开数据库，建立连接
	 db,_ := sql.Open("mysql","root:hanru1314@tcp(127.0.0.1:3306)/my1802?charset=utf8")

	 //stpt3：查询数据库
	 rows,err:=db.Query("SELECT empno,ename,job,hiredate,sal,deptno FROM emp WHERE deptno=?",30)
	fmt.Println("--->",err.Error())
	 fmt.Println(rows.Columns()) //[empno ename job hiredate sal deptno]
	 //思路一：定义一个map，用于存储从数据库中查询出来的数据，字段作为key，string，数据作为value，任意类型，空接口
	 map1:=make(map[string]interface{})
	 datas := make([] map[string]interface{},0)

	 //思路二：创建slice，存入struct，
	 datas2:=make([] Emp,0)
	 //step4：操作结果集获取数据
	 for rows.Next(){
	 	var empno int
	 	var ename string
	 	var job string
	 	var hiredate string
	 	var sal float64
	 	var deptno int
		if err:=rows.Scan(&empno,&ename,&job,&hiredate,&sal,&deptno);err!=nil{
			fmt.Println("获取失败。。")
		}
		//fmt.Println(empno,ename,job,hiredate,sal,deptno)
		//将读取到的数据，存入了map中
		map1["empno"]=empno
		map1["ename"]=ename
		map1["job"]=job
		map1["hiredate"]=hiredate
		map1["sal"]=sal
		map1["deptno"]=deptno
		//将map存入切片中
		datas = append(datas,map1)


		//思路二：每读取一行，创建一个emp对象，存入datas2中
		emp := Emp{empno,ename,job,hiredate,sal,deptno}
		datas2 =append(datas2, emp)
	 }
	 //step5：关闭资源
	 rows.Close()
	 db.Close()

	 //遍历切片
	 //fmt.Println("empno\tename\tjob\thiredate\tsal\tdeptno")
	 //for _,v:=range datas{
	 //	for _,val:=range v{
	 //		fmt.Print(val,"\t")
		//}
		//fmt.Println()
	 //
	 //}

	 for _,v:=range  datas2{
	 	fmt.Println(v)
	 }


}
/*
查询：处理查询后的结果：
	思路一：将数据，存入到map中
	思路二：创建结构体：
 */
