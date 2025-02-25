package models

import "time"

type Role struct{
	Name string `json:"role_name"`
	TimeToEnd time.Time
}

type User struct{
	ID uint64 `json:"id"`
	Token string `json:"token"`
	Firstname string `json:"firstname"`
	Secondname string `json:"secondname"`
	Role Role`json:"role"`
}

type Document struct{
	ID uint64 `json:"id"`
	About string `json:"about"`
	CreatedAt time.Time `json:"created_at"`
	Content string `json:"content"`
	Roles []string  `json:"roles"`
}