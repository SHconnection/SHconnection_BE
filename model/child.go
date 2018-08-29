package model

type Child struct {
	BasicModel
	Name 		string 	`json: "name" gorm: "type:varchar(20)"`

	//  child 属于 class
	ClassID 	uint     // foreign key

	Parents		[]Person

	Sid  		string `json: "sid" gorm: "type:varchar(20)"`
}
