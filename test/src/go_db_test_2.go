package main

import (
	"database/sql"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
)

const (
	DB_NAME = "parkdb"
	DB_USER = "root"
	DB_PASS = "root123"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Alias string `json:"alias"`
}

func OpenDB() *sql.DB {
	db, err := sql.Open("mymysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	if err != nil {
		panic(err)
	}
	return db
}

func UserById(id int) User {
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT `uid`, `username`,`password` FROM `user` WHERE uid=?", id)
	user := new(User)
	row.Scan(&user.Id, &user.Name, &user.Alias)
	return user
}
