package main

import (
	"fmt"
	"database/sql"
	"github.com/ziutek/mymysql/godrv"
)

type User struct {
	uid      int `PK`
	username string
	password string
}

func main() {
	// 设置连接编码
	godrv.Register("SET NAMES utf8")

	// 连接数据库
	db, err := sql.Open("mymysql", "tcp:127.0.0.1:3306*parkdb/root/root123")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 插入数据
	stmt, err := db.Prepare("insert into user values(null, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// sql参数
	result, err := stmt.Exec("jeff", "mu.jm")
	if err != nil {
		panic(err)
	}

	// 获取影响的行数
	affect, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", affect)

	// 获取自增id
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", id)

	// 查询数据
	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// 获取的用户
	users := []User{}

	// 读取数据
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.uid, &user.username, &user.password)
		if nil != err {
			panic(err)
		}

		users = append(users, user)
	}

	// 显示用户信息
	for _, user := range users {
		fmt.Printf("%d, %s, %s\n", user.uid, user.username, user.password)
	}
}
