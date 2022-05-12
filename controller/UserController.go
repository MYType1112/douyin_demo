package controller

import (
	"douyin-demo/dao"
	"douyin-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserLoginResponse struct {
	dao.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	dao.Response
	User dao.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//Map<String, Object> map = userService.register(user);
	registerMap := service.Register(username, password)
	if _, ok := registerMap["userId"]; ok {
		userId, _ := strconv.ParseInt(registerMap["userId"], 10, 64)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: dao.Response{StatusCode: 0},
			UserId:   userId,
			Token:    registerMap["Token"],
		})
	} else {
		var msg string
		if val, ok := registerMap["usernameMsg"]; ok {
			msg = val
		}
		if val, ok := registerMap["passwordMsg"]; ok {
			msg = msg + val
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: msg},
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	loginMap := service.Login(username, password)

	if _, ok := loginMap["userId"]; ok {
		userId, _ := strconv.ParseInt(loginMap["userId"], 10, 64)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: dao.Response{StatusCode: 0},
			UserId:   userId,
			Token:    loginMap["ticket"],
		})
	} else {
		var msg string
		if val, ok := loginMap["usernameMsg"]; ok {
			msg = val
		}
		if val, ok := loginMap["passwordMsg"]; ok {
			msg = msg + val
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: msg},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	infoMap, user := service.UserInfo(token)

	if _, ok := infoMap["userId"]; ok {
		c.JSON(http.StatusOK, UserResponse{
			Response: dao.Response{StatusCode: 0},
			User:     *user,
		})
	} else {
		var msg string
		if val, ok := infoMap["errMsg"]; ok {
			msg = val
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: dao.Response{StatusCode: 1, StatusMsg: msg},
		})
	}
}
