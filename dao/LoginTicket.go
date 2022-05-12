package dao

import "time"

type LoginTicket struct {
	Id      int64     `json:"id,omitempty"`
	UserId  int64     `json:"username,omitempty"`
	Ticket  string    `json:"password,omitempty"`
	Status  int       `json:"salt,omitempty"`
	Expired time.Time `json:"email,omitempty"`
}

func (l *LoginTicket) SetUserId(userId int64) {
	l.UserId = userId
}
func (l *LoginTicket) SetTicket(ticket string) {
	l.Ticket = ticket
}
func (l *LoginTicket) SetStatus(status int) {
	l.Status = status
}
func (l *LoginTicket) SetExpired(expired time.Time) {
	l.Expired = expired
}
