package dao

import "time"

type User struct {
	Id             int64     `json:"id,omitempty"`
	Username       string    `json:"username,omitempty"`
	Password       string    `json:"password,omitempty"`
	Salt           string    `json:"salt,omitempty"`
	Email          string    `json:"email,omitempty"`
	Type           int32     `json:"type,omitempty"`
	Status         int32     `json:"status,omitempty"`
	ActivationCode string    `json:"activationCode,omitempty"`
	HeaderUrl      string    `json:"headerUrl,omitempty"`
	CreateTime     time.Time `json:"createTime,omitempty"`
}

func (u *User) SetUsername(name string) {
	u.Username = name
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetSalt(salt string) {
	u.Salt = salt
}

func (u *User) GetSalt() string {
	return u.Salt
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetType(usertype int32) {
	u.Type = usertype
}

func (u *User) GetType() int32 {
	return u.Type
}

func (u *User) SetStatus(status int32) {
	u.Status = status
}

func (u *User) GetStatus() int32 {
	return u.Status
}

func (u *User) SetActivationCode(activationCode string) {
	u.ActivationCode = activationCode
}

func (u *User) GetActivationCode() string {
	return u.ActivationCode
}

func (u *User) SetHeaderUrl(headerUrl string) {
	u.HeaderUrl = headerUrl
}

func (u *User) GetHeaderUrl() string {
	return u.HeaderUrl
}

func (u *User) SetCreateTime(createTime time.Time) {
	u.CreateTime = createTime
}

func (u *User) GetCreateTime() time.Time {
	return u.CreateTime
}
