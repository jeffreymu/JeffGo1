package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	insert()
}

func insert() {
	db, err := sql.Open("mysql", "root:@/parkdb?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`INSERT user (uid,username,password) values (?,?,?)`)
	//checkErr(err)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	res, err := stmt.Exec("jeffrey", "mumu")
	//checkErr(err)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	id, err := res.LastInsertId()
	//checkErr(err)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println(id)
}
