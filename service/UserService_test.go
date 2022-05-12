package service

import (
	"douyin-demo/dao"
	"fmt"
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := dao.InitDB()
	if err != nil {
		fmt.Printf("err:%#v\n", err)
	} else {
		println("connect success")
	}
	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	err := dao.InitDB()
	if err != nil {
		fmt.Printf("err:%#v\n", err)
	} else {
		println("connect success")
	}
	register := Register("mytype1", "123456")
	assert.NotEqual(t, nil, register)
	for key, value := range register {
		fmt.Println(key, "----->", value)
	}
}

func TestLogin(t *testing.T) {
	err := dao.InitDB()
	if err != nil {
		fmt.Printf("err:%#v\n", err)
	} else {
		println("connect success")
	}
	login := Login("mytype1", "123456")
	assert.NotEqual(t, nil, login)
	for key, value := range login {
		fmt.Println(key, "----->", value)
	}
}

func TestLogout(t *testing.T) {
	err := dao.InitDB()
	if err != nil {
		fmt.Printf("err:%#v\n", err)
	} else {
		println("connect success")
	}
	logout := Logout("215ae1a007124ac19b8aeed464364be1")
	println(logout)
}

func TestUserInfo(t *testing.T) {
	err := dao.InitDB()
	if err != nil {
		fmt.Printf("err:%#v\n", err)
	} else {
		println("connect success")
	}
	info, _ := UserInfo("215ae1a007124ac19b8aeed464364be1")
	assert.NotEqual(t, nil, info)
	for key, value := range info {
		fmt.Println(key, "----->", value)
	}
}
