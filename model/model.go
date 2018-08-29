package model

import "time"

type BasicModel struct {
	ID 				uint `gorm:"primary_key"`
	CreatedAt 		time.Time
	UpdatedAt   	time.Time
}

type Person struct {

	Name  			string `json: "name" gorm:"type:varchar(20)"`
	// 区分老师还是家长
	Type  			string `json: "type" gorm: "type:varchar(20)"`
	// 电话号码
	Tel  			string `json: "tel" gorm: "type:varchar(20)"`
	// password hash
	PasswordHash 	string `json: "-" gorm: gorm: "type:varchar(200)"`
	Wechat 			string `json: "wechat" gorm: "type:varchar(30)"`
	Intro 			string `json: "intro" gorm: "type:varchar(300)"`
	Avatar 			string `json: "avatar" gorm: "type:varchar(200)"`
} 
