package model

type Child struct {
	BasicModel
	Name 		string 	`json: "name" gorm: "type:varchar(20)"`
	TheClass 		Class
	Sid  		string `json: "sid" gorm: "type:varchar(20)"`
	Avatar   	string `json: "avatar" gorm: "type: vrachar(100)"`
}
