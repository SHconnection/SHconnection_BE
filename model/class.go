package model

type Class struct {
	BasicModel
	ClassName 		string 	 `json: "classname" gorm: "unique"`

	MainTeacher 	Teacher  //  一个 class 包含一个 班主任

	// 班级的老师，多对多关系
	Teachers        []Teacher    `gorm:"many2many:class_teachers;"`

	// 班级的孩子，一对多关系
	Childs          []Child
}

