package models

import "gorm.io/gorm"

type Ingredient struct {
	ID				uint			`gorm:"primaryKey"`
	Quantity		uint			`json:"quantity"`
	Unit			string			`json:"unit"`
	Name			string			`json:"name"`
	RecipieRefer	uint
}

type Step struct {
	ID				uint			`gorm:"primaryKey"`
	Step			string			`json:"step"`
	StepsRefer		uint
}

type Dish struct {
	ID				uint			`gorm:"primaryKey"`
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
