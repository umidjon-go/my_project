package models

import "time"

type User struct {
	Id          int        `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	BranchId    int        `json:"branch_id"`
	Phone       string     `json:"phone"`
	Status      string     `json:"status"`
	RatingScore float64    `json:"rating_score"`
	Role        string     `json:"role"`
	Position    string     `json:"position"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
