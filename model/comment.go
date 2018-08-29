package model

type Comment struct {
	BasicModel
	Content 		string 	`json:"content"`
	Likes 			uint   	`json: "-"`

	// Feed ID foreign key
	FeedID          uint

	// 一对多关系, foreign key
	Person			Person
	PersonID		uint
}
