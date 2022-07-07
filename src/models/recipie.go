package models

import "gorm.io/gorm"

type ingredient struct {
	Quantity	int32		`json:"quantity"`
	Unit		string		`json:"unit"`
	Name		string		`json:"name"`
}

type dish struct {
	Dish	string		`json:"dish"`
	Steps	[]string	`json:"steps"`
}

type Recipie struct {
	gorm.Model
	Title			string			`json:"title"`
	Author			string 			`json:"author"`
	Ingredients		[]ingredient	`json:"ingredients"`
	Instructions	[]dish			`json:"instructions"`
}
