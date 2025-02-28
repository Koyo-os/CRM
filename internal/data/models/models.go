package models

import "time"

const TIME_LAYOUT = "02.01.2006"

type Role struct{
	Name string `json:"role_name"`
	RunningOut bool `json:"running_out" bson:"running_out"`
	TimeToEnd time.Time `json:"time_to_end" bson:"time_to_end"`
	TypeRole [3]rune `json:"type_role" bson:"type_role"`
	CanAddDoc bool `json:"can_add_doc" bson:"can_add_doc"`
}

type DocRole struct{
	Name string `json:"role_name"`
	TypeRole [3]rune `json:"type_role" bson:"type_role"`
}

type User struct{
	ID uint64 `json:"id" bson:"id"`
	Key string `json:"key" bson:"key"`
	Firstname string `json:"firstname" bson:"firstname"`
	Secondname string `json:"secondname" bson:"secondname"`
	Role []Role`json:"role" bson:"role"` 
	SuperUser bool `json:"super_user" bson:"super_user"`
}

type Document struct{
	ID uint64 `json:"id" bson:"id"`
	About string `json:"about" bson:"about"` 
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Content string `json:"content" bson:"content"`
	Roles []string  `json:"roles" bson:"roles"`
}