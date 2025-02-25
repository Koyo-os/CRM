package models

import "time"

type Role struct{
	Name string `json:"role_name"`
	TimeToEnd time.Time
}

type User struct{
	ID uint64 `json:"id" bson:"id"`
	Key string `json:"key" bson:"key"`
	Firstname string `json:"firstname" bson:"firstname"`
	Secondname string `json:"secondname" bson:"secondname"`
	Role Role`json:"role" bson:"role"` 
	SuperUser bool `json:"super_user" bson:"super_user"`
}

type Document struct{
	ID uint64 `json:"id" bson:"id"`
	About string `json:"about" bson:"about"` 
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Content string `json:"content" bson:"content"`
	Roles []string  `json:"roles" bson:"roles"`
}