package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Quantity		int				`json:"quantity"`
	Unit			string			`json:"unit"`
	Name			string			`json:"name"`
	RecipieRefer	uint
}

type Step struct {
	gorm.Model
	Step			string			`json:"step"`
	StepsRefer		uint
}

type Dish struct {
	gorm.Model
	Dish			string			`json:"dish"`
	Steps			[]Step			`gorm:"foreignKey:StepsRefer" json:"steps"`
	RecipieRefer	uint
}

type Recipie struct {
	gorm.Model
	Title			string			`json:"title"`
	Author			string 			`json:"author"`
	Ingredients		[]Ingredient	`gorm:"foreignKey:RecipieRefer" json:"ingredients"`
	Instructions	[]Dish			`gorm:"foreignKey:RecipieRefer" json:"instructions"`
}
