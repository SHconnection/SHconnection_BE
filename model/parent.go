package model


type Parent struct {
	BasicModel
	Name 			string `json: "name" gorm: "type:varchar(20)"`
	TheChild 			Child
	ChildRelation 	string `json: child_relation gorm: "type:varchar(30)"`
	Tel 			string `json: tel gorm: "type:varchar(20)"`
	PasswordHash    string `json: "-" gorm: "type:varchar(200)"`
	Intro 			string `json: "-" gorm: "type:varchar(300)"`
	Avatar 			string `json: "avatar" gorm: "type:varchar(200)"`
}