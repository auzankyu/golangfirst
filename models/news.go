package models

type News struct {
	Date     string `json:"date"`
	Content  string `json:"content"`
	Category_id string `json:"category_id"`
}
