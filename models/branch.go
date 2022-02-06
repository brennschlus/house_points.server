package models

type Branch struct {
	ID          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Information string `json:"information"`
}
