package model


type Parent struct {
	BasicModel
	Person

	ChildID         uint
	//与孩子的关系
	ChildRelation 	string `json: child_relation gorm: "type:varchar(30)"`

}