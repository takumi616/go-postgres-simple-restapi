package entity

type Vocabulary struct {
	Id        int       `json:"id"`
	Title     string	`json:"title"`
	Sentence  string	`json:"sentence"`
	Meaning   string	`json:"meaning"`
}
