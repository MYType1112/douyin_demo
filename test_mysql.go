package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

type User struct {
	Id             int64  `json:"id,omitempty"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	Salt           string `json:"salt,omitempty"`
	Email          string `json:"email,omitempty"`
	Type           int64  `json:"type,omitempty"`
	Status         int64  `json:"status,omitempty"`
	ActivationCode string `json:"activationCode,omitempty"`
	HeaderUrl      string `json:"headerUrl,omitempty"`
	CreateTime     string `json:"createTime,omitempty"`
}

var db *sql.DB

func initDB() (err error) {
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

func queryOneRow() {
	s := "select * from user where id = ?"
	var u User
	err := db.QueryRow(s, 146).Scan(&u.Id, &u.Username, &u.Password, &u.Salt,
		&u.Email, &u.Type, &u.Status, &u.ActivationCode, &u.HeaderUrl, &u.CreateTime)
	if err != nil {
		fmt.Printf("scan failed, err :%v\n", err)
		return
	}
	fmt.Printf("name: %s, status: %v\n", u.Username, u.Status)
}

func queryOneRow2() {
	s := "select * from user where username = ?"
	var u User
	err := db.QueryRow(s, "lhh").Scan(&u.Id, &u.Username, &u.Password, &u.Salt,
		&u.Email, &u.Type, &u.Status, &u.ActivationCode, &u.HeaderUrl, &u.CreateTime)
	if err != nil {
		fmt.Printf("scan failed, err :%v\n", err)
		return
	}
	fmt.Printf("name: %s, status: %v\n", u.Username, u.Status)
}

func updateData() {
	sql := "update user set status = ? where id = ?"
	ret, err := db.Exec(sql, "1", "146")
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("update rows failed, err: %v\n", err)
		return
	}
	fmt.Printf("update success, update row: %v\n", rows)
}

func insert() {
	sqlStr := "insert into user (username, password, salt, email, type,status,activation_code,header_url,create_time) " +
		"values(?,?,?,?,?,?,?,?,?)"
	ret, err := db.Exec(sqlStr, "user2", "123456", "49f10", "1234567@qq.com", 1, 1, nil, nil, time.Now())
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId error: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id: %v\n", theID)
}

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomInt(n int) int {
	maxNum := n
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(maxNum)
	return randNumber
}

//func main() {
//
//	randomString := fmt.Sprintf("http://images.nowcoder.com/head/%dt.png", GetRandomInt(1000))
//	fmt.Printf("%v\n", randomString)
//}
