package models

import "time"

type Diary struct {
	Id        uint       `json:"id"`
	Uid       string     `json:"uid"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"create_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `json:"delete_at"`
}
