package models

import "time"

type Account struct {
	Id        uint      `json:"id"`
	Uid       string    `json:"uid" gorm:"unique"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  []byte    `json:"-"`
	CreatedAt time.Time `json:"create_at"`
}
