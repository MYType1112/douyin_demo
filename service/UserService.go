package service

import (
	"douyin-demo/dao"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Register(username string, password string) map[string]string {
	m := make(map[string]string, 2)
	if dao.IsBlank(username) {
		m["usernameMsg"] = "账号格式错误!"
		return m
	}
	if dao.IsBlank(password) {
		m["passwordMsg"] = "密码格式错误!"
		return m
	}
	userDao := dao.NewUserDaoInstance()
	user := userDao.SelectByName(username)
	//user := dao.NewUserDaoInstance().SelectByName(username)
	if user != nil {
		m["usernameMsg"] = "该账号已存在!"
		return m
	} else {
		//var user *dao.User
		var u dao.User
		user = &u
		user.SetUsername(username)
		user.SetSalt(dao.GetRandomString(5))
		user.SetPassword(dao.Md5(password + user.GetSalt()))
		user.SetType(0)
		user.SetStatus(1)
		user.SetActivationCode(dao.GetRandomString(32))
		user.SetHeaderUrl(fmt.Sprintf("http://images.nowcoder.com/head/%dt.png", dao.GetRandomInt(1000)))
		user.SetCreateTime(time.Now())
		theId := userDao.InsertUser(user)
		m["userId"] = strconv.FormatInt(theId, 10)
		m["Password"] = user.Password
		return m
	}
}

func Login(username string, password string) map[string]string {
	m := make(map[string]string, 2)
	if dao.IsBlank(username) {
		m["usernameMsg"] = "账号格式错误!"
		return m
	}
	if dao.IsBlank(password) {
		m["passwordMsg"] = "密码格式错误!"
		return m
	}
	userDao := dao.NewUserDaoInstance()
	user := userDao.SelectByName(username)
	//user := dao.NewUserDaoInstance().SelectByName(username)
	if user == nil {
		m["usernameMsg"] = "该账号不存在!"
		return m
	}
	// 验证状态
	if user.GetStatus() == 0 {
		m["usernameMsg"] = "该账号未激活!"
		return m
	}
	// 验证密码
	password = dao.Md5(password + user.GetSalt())
	if strings.Compare(user.GetPassword(), password) != 0 {
		m["passwordMsg"] = "密码不正确!"
		return m
	}
	// 生成登录凭证
	loginDao := dao.NewLoginTicketDaoInstance()
	var l dao.LoginTicket
	lo := &l
	lo.SetUserId(user.Id)
	lo.SetTicket(dao.GetUUID())
	lo.SetStatus(0)
	lo.SetExpired(time.Now().Add(time.Hour * 24 * 30))
	loginDao.InsertLoginTicket(lo)

	m["userId"] = strconv.FormatInt(user.Id, 10)
	m["ticket"] = lo.Ticket
	return m
}

func Logout(ticket string) int {
	return dao.UpdateLoginStatus(1, ticket)
}

func UserInfo(token string) (map[string]string, *dao.LoginTicket) {
	m := make(map[string]string, 2)
	if dao.IsBlank(token) {
		m["errMsg"] = "参数错误!"
		return m, nil
	}
	loginTicketDao := dao.NewLoginTicketDaoInstance()
	loginTicket := loginTicketDao.SelectByTicket(token)

	if loginTicket == nil {
		m["errMsg"] = "用户未登录!"
		return m, nil
	}
	if loginTicket.Status == 1 {
		m["errMsg"] = "用户未登录!"
		return m, nil
	}
	if loginTicket.Expired.Before(time.Now()) {
		m["errMsg"] = "用户登录过期!"
		return m, nil
	}
	userDao := dao.NewUserDaoInstance()
	user := userDao.SelectById(loginTicket.UserId)
	if user == nil {
		m["errMsg"] = "用户不存在!"
		return m, nil
	}

	m["userId"] = strconv.FormatInt(loginTicket.UserId, 10)
	m["Token"] = loginTicket.Ticket
	return m, loginTicket
}
