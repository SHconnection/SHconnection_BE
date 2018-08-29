package model

type Teacher struct {
	BasicModel
	Name  			string `json: "name" gorm:"type:varchar(20)"`
	Type  			string `json: "type" gorm: "type:varchar(20)"`
	Tel  			string `json: "tel" gorm: "type:varchar(20)"`
	PasswordHash 	string `json: "-" gorm: gorm: "type:varchar(200)"`
	Wechat 			string `json: "wechat" gorm: "type:varchar(30)"`
	Intro 			string `json: "intro" gorm: "type:varchar(300)"`
	Avatar 			string `json: "avatar"`
}
