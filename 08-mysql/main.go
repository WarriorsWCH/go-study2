

package main 

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
// https://github.com/go-sql-driver/mysql

func main(){
	// 终端创建数据库---建库：create database Mytest；（Mytest库名）
	// 创建表：
	// CREATE TABLE `my_info` (`uid` INT(10) NOT NULL AUTO_INCREMENT,
	// `username` VARCHAR(64) NULL DEFAULT NULL,
	// `departname` VARCHAR(64) NULL DEFAULT NULL,
	// `created` DATE NULL DEFAULT NULL,PRIMARY KEY (`uid`));
	// 打开某个数据库（比如数据库：Mytest)：use Mytest;
	// 显示本库中的所有表：show tables;
	// 建表：create table 表名 (字段设定列表)； // :creat table mytest_acount (col1 INT, col2 CHAR(5), col3 DATE);  表至少一列。
	db, err := sql.Open("mysql", "root:wang1234@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
    // 插入数据
    stmt, err := db.Prepare("INSERT my_info SET username=?,departname=?,created=?")
    checkErr(err)
    res, err := stmt.Exec("test", " 研发部门", "2017-12-09")
    checkErr(err)
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(id)
	
    // 更新数据
    stmt, err = db.Prepare("update my_info set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("test", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	
	// 查询数据
	rows, err := db.Query("SELECT * FROM my_info")
	checkErr(err)
	for rows.Next() {
		 var uid int
		 var username string
		 var department string
		 var created string
		 err = rows.Scan(&uid, &username, &department, &created)
		 checkErr(err)
		 fmt.Println(uid)
		 fmt.Println(username)
		 fmt.Println(department)
		 fmt.Println(created)
	}
	
	// 删除数据
	stmt, err = db.Prepare("delete from my_info where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
    checkErr(err)
    affect, err = res.RowsAffected()
    checkErr(err)
    fmt.Println(affect)

    defer db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}















