package model

type Comment struct {
	BasicModel
	CommentFeed 	Feed
	Content 		string 	`json:"content"`
	Likes 			uint   	`json: "-"`
}
