package data

type Student struct {
	Id        int    `gorm:"primaryKey"`
	Name      string `json:"name"`
	Number    string `json:"number"`
	Graduated bool   `json:"graduated"`
}

var Students []Student
