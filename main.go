package main

import (
	"douyin-demo/dao"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	err := dao.InitDB()
	if err != nil {
		fmt.Printf("err:%#v\n", err)
	} else {
		println("connect success")
	}
}

func main() {
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
