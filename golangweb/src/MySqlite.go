package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
	"net/smtp"
)

var err error

func main() {
	db, err := sql.Open("sqlite3", "foo.db")
	defer db.Close()
	checkErr(err)
	err = createTable(db)
	if err != nil {
		return
	}
	//插入数据
	//stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)
	defer stmt.Close()
	res, err := stmt.Exec("astaxie", "研发部门", time.Now().Format("2006-01-02"))
	checkErr(err)
	id, err := res.LastInsertId()
	fmt.Println("LastInsertId-->", id)
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
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
	//删除数据
	//	stmt, err = db.Prepare("delete from userinfo where uid=?")
	//	checkErr(err)
	//
	//	res, err = stmt.Exec(id)
	//	checkErr(err)
	//
	//	affect, err = res.RowsAffected()
	//	checkErr(err)

	fmt.Println(affect)



}

/***
*create table
 */
func createTable(db *DB) error {
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS userinfo (uid INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, departname TEXT,created Text);")
	if err != nil {
		log.Fatalln("could not create table:", err)
	}
	return err
}
func dropTable(db *DB) {
	//Drop, create and insert data into the test table
	_, err = db.Exec("DROP TABLE IF EXISTS userinfo;")
	if err != nil {
		log.Fatalln("could not drop table:", err)
	}

}
func checkErr(err error) {
	if err != nil {
		("error-->>:", err)
	}
}
