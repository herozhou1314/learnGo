package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {

	db, err := sql.Open("sqlite3", "locked.sqlite")
	if err != nil {
		log.Fatalln("could not open database:", err)
	}
	defer db.Close()

	// Drop, create and insert data into the test table
	_, err = db.Exec("DROP TABLE IF EXISTS src;")
	if err != nil {
		log.Fatalln("could not drop table:", err)
	}
	_, err = db.Exec("DROP TABLE IF EXISTS dst;")
	if err != nil {
		log.Fatalln("could not drop table:", err)
	}
	_, err = db.Exec("CREATE TABLE src (id INTEGER PRIMARY KEY AUTOINCREMENT, data1 INTEGER, data2 INTEGER);")
	if err != nil {
		log.Fatalln("could not create table:", err)
	}
	_, err = db.Exec("CREATE TABLE dst (id INTEGER, data1 INTEGER, data2 INTEGER);")
	if err != nil {
		log.Fatalln("could not create table:", err)
	}

	for i := 0; i < 100; i++ {
		//_, err = db.Exec("INSERT INTO src (id, data1, data2) VALUES (?, ?, ?);", i, 100+i, 1000+i)
		_, err = db.Exec("INSERT INTO src ( data1, data2) VALUES ( ?, ?);",  100+i, 1000+i)
		if err != nil {
			log.Fatalln("could not insert into table:", err)
		}
	}

	rows, err := db.Query("SELECT id, data1, data2 FROM src;")
	if err != nil {
		log.Fatalln("could not select data:", err)
	}
	defer rows.Close()

	insert, err := db.Prepare("INSERT INTO dst (id, data1, data2) VALUES (?, ?, ?);")
	if err != nil {
		log.Fatalln("could not prepare statement:", err)
	}
	defer insert.Close()

	for rows.Next() {
		var id, data1, data2 int64

		err = rows.Scan(&id, &data1, &data2)
		if err != nil {
			log.Fatalln("could not scan row:", err)
		}

		_, err = insert.Exec(id, data1, data2)
		if err != nil {
			log.Fatalln("could not insert data into dst:", err)
		}
	}
	insert.Close()
	rows.Close()

	return
}
