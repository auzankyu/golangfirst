package models

type ResponseSource struct {
	Status string `json:"status"`
	Sources []Source `json:"sources"`
}

type Source struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Url string `json:"url"`
	Category string `json:"category"`
	Language string `json:"language"`
	Country string `json:"country"`
}
