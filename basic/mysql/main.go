package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var DB *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test3")
	if err != nil {
		panic(err)
	}
	DB = database
}

func main() {
	//insert()
	update()
	query()
	transaction()
}

func transaction() {
	conn, err := DB.Begin()
	if err != nil {
		panic(err)
	}

	r, err := conn.Exec("insert into person(username, sex, email) values(?,?,?)", "stu001", "man", "stu01@163.com")
	if err != nil {
		conn.Rollback()
		panic(err)
	}
	r, err = conn.Exec("insert into person(username, sex, email) values(?,?,?)", "stu001", "man", "stu01@163.com")
	if err != nil {
		conn.Rollback()
		panic(err)
	}
	conn.Commit()
	fmt.Println(r)
}

func insert() {
	r, err := DB.Exec("insert into person(username, sex, email) values(?,?,?)", "stu001", "man", "stu01@163.com")
	if err != nil {
		panic(err)
	}
	id, err := r.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("insert succ:", id)
}
func query() {
	var person []Person
	err := DB.Select(&person, "select user_id,username,sex,email from person where user_id=?", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("insert succ:", person)
}

func update() {
	_, err := DB.Exec("update person set username=? where user_id=?", "stu0001", 1)
	if err != nil {
		panic(err)
	}
}
