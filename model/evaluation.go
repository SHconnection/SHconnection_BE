package model

import "time"

type Evaluation struct {
	BasicModel
	Content  	string `json: "content" gorm: "type:text"`
	Time 		time.Time
	Type 		string `json: "type" gorm: "type:varchar(20)"`  // 来自老师还是家长
	// 先判断来自家长还是老师，来自家长就设置家长 来自老师就设置老师，另一个为nil
	FromTeacher		Teacher
	FromParent 		Parent
	// 多设置一些，以备用
	Score1 		float64 `json: "score1" gorm: "type:float"`
	Score2 		float64 `json: "score2" gorm: "type:float"`
	Score3 		float64 `json: "score3" gorm: "type:float"`
	Score4 		float64 `json: "score4" gorm: "type:float"`
	Score5 		float64 `json: "score5" gorm: "type:float"`
	Score6 		float64 `json: "score6" gorm: "type:float"`
	Score7 		float64 `json: "score7" gorm: "type:float"`
	Score8 		float64 `json: "score8" gorm: "type:float"`
	Score9 		float64 `json: "score9" gorm: "type:float"`
	Score10 		float64 `json: "score10" gorm: "type:float"`
	Score11 		float64 `json: "score11" gorm: "type:float"`
	Score12 		float64 `json: "score12" gorm: "type:float"`
	Score13 		float64 `json: "score13" gorm: "type:float"`
	Score14 		float64 `json: "score14" gorm: "type:float"`
	Score15 		float64 `json: "score15" gorm: "type:float"`
	Score16 		float64 `json: "score16" gorm: "type:float"`
	Score17 		float64 `json: "score17" gorm: "type:float"`
	Score18 		float64 `json: "score18" gorm: "type:float"`
}
