package models

import "time"

type Payment struct {
	Id          int        `json:"id"`
	CourseId    int        `json:"course_id"`
	StudentId   int        `json:"student_id"`
	GroupId     int        `json:"group_id"`
	UserId      int        `json:"user_id"`
	Amount      int64      `json:"amount"`
	PaymentType string     `json:"payment_type"`
	Status      string     `json:"status"`
	PaidAt      *time.Time `json:"paid_at"`
	Comment     string     `json:"comment"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
