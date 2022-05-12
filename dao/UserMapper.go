package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func QueryById(id int64) *User {
	s := "select * from user where id = ?"
	var u User
	err := db.QueryRow(s, id).Scan(&u.Id, &u.Username, &u.Password, &u.Salt,
		&u.Email, &u.Type, &u.Status, &u.ActivationCode, &u.HeaderUrl, &u.CreateTime)
	if err != nil {
		fmt.Printf("scan failed, err :%v\n", err)
		return nil
	}
	fmt.Printf("name: %s, status: %v\n", u.Username, u.Status)
	return &u
}

func QueryByName(name string) *User {
	s := "select * from user where username = ?"
	var u User
	err := db.QueryRow(s, name).Scan(&u.Id, &u.Username, &u.Password, &u.Salt,
		&u.Email, &u.Type, &u.Status, &u.ActivationCode, &u.HeaderUrl, &u.CreateTime)
	if err != nil {
		fmt.Printf("scan failed, err :%v\n", err)
		return nil
	}
	fmt.Printf("name: %s, status: %v\n", u.Username, u.Status)
	return &u
}

func QueryByToken(token string) *User {
	s := "select * from user where password = ?"
	var u User
	err := db.QueryRow(s, token).Scan(&u.Id, &u.Username, &u.Password, &u.Salt,
		&u.Email, &u.Type, &u.Status, &u.ActivationCode, &u.HeaderUrl, &u.CreateTime)
	if err != nil {
		fmt.Printf("scan failed, err :%v\n", err)
		return nil
	}
	fmt.Printf("name: %s, status: %v\n", u.Username, u.Status)
	return &u
}

func UpdateStatus(id int64, status int64) int {
	s := "update user set status = ? where id = ?"
	ret, err := db.Exec(s, status, id)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return 0
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("update rows failed, err: %v\n", err)
		return 0
	}
	fmt.Printf("update success, update row: %v\n", rows)
	return 1
}

func UpdatePassword(id int64, password string) int {
	s := "update user set password = ? where id = ?"
	ret, err := db.Exec(s, password, id)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return 0
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("update rows failed, err: %v\n", err)
		return 0
	}
	fmt.Printf("update success, update row: %v\n", rows)
	return 1
}

func InsertUser(user *User) int64 {
	sqlStr := "insert into user (username,password,salt,email,type,status,activation_code,header_url,create_time) " +
		"values(?,?,?,?,?,?,?,?,?)"
	ret, err := db.Exec(sqlStr, user.Username, user.Password, user.Salt,
		user.Email, user.Type, user.Status, user.ActivationCode, user.HeaderUrl, user.CreateTime)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return 0
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId error: %v\n", err)
		return 0
	}
	fmt.Printf("insert success, the id: %v\n", theID)
	return theID
}
