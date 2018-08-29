package model

type Feed struct {
	BasicModel
	FeedClass		Class
	PostTeacher		Teacher
	Type 			string `json: "type" gorm: "type:varchar(20)"`
	Content 		string `json: "content" gorm: "type:text"`
	Likes			uint   `json: "-"`
}
