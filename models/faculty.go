package models

type Faculty struct {
	BranchID string `json:"branchID" `
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name"`
	Score    uint   `json:"score"`
}
