package models

import "time"

type Course struct {
	Id          int        `json:"id"`
	UserId      int        `json:"user_id"`
	BranchId    int        `json:"branch_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
