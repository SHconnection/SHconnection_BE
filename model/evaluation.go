package model

type Evaluation struct {
	BasicModel
	Content  		string `json: "content" gorm: "type:text"`

	Person     		Person
	PersonID        uint

	Child			Child
	ChildID         uint

	Scores			[]Score
}

type Score struct {
	BasicModel
	Score			float64	`json: "score" gorm: "type:float"`
	Name			string	`json: "name" gorm: "type:varchar(20)"`
	// foreign key 
	EvaluationID 	uint
}

