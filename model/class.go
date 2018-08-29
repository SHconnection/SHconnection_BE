package model

type Class struct {
	BasicModel
	ClassName 		string 	 `json: "classname" gorm: "unique"`
	MainTeacher 	Teacher  // 外键 班主任
}

type Teacher2Class struct {
	BasicModel
	TheTeacher  		Teacher   // 外键
	TheClass 			Class     // 外键
}