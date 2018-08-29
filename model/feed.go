package model

type Feed struct {
	BasicModel

	// 一对多关系，foreign key
	Class			Class
	ClassID     	uint

	// 一对多关系，foreign key
	Teacher			Teacher
	TeacherID   	uint

	Type 			string `json: "type" gorm: "type:varchar(20)"`
	Content 		string `json: "content" gorm: "type:text"`
	Likes			uint   `json: "-"`

	// feed的评价 一对多关系
	Comments        []Comment
}
