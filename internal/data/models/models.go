package models

import "time"

type User struct{
	ID uint64
	Token string
	Firstname string
	Secondname string
	Role string
}

type Document struct{
	ID uint64 `json:"id"`
	About string `json:"about"`
	CreatedAt time.Time `json:"created_at"`
	Content string `json:"content"`
	Roles []string  `json:"roles"`
}