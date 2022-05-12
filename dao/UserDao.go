package dao

import "sync"

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}
func (*UserDao) SelectById(id int64) *User {
	return QueryById(id)
}

func (*UserDao) SelectByName(username string) *User {
	return QueryByName(username)
}

func (*UserDao) SelectByToken(token string) *User {
	return QueryByToken(token)
}

func (*UserDao) InsertUser(user *User) int64 {
	return InsertUser(user)
}

func (*UserDao) UpdateStatus(id int64, status int64) int {
	return UpdateStatus(id, status)
}

func (*UserDao) UpdatePassword(id int64, password string) int {
	return UpdatePassword(id, password)
}

//func (*UserDao) updateHeader(id int64, headerUrl string) bool {
//	return true
//}
//func (*UserDao) selectByEmail(email string) *User {
//	return nil
//}
