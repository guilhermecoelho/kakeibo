package models

type Expense struct {
	Id         string   `json:"id"  gorm:"primary_key"`
	Name       string   `json:"name"`
	Date       string   `json:"date"`
	CategoryId int      `json:"categoryId"`
	Amount     float32  `json:"amount"`
	Note       string   `json:"note"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}
