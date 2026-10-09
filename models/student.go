package models

import "time"

type Student struct {
	Id          int        `json:"id"`
	CourseId    int        `json:"course_id"`
	UserId      int        `json:"user_id"`
	GroupId     int        `json:"group_id"`
	FullName    string     `json:"full_name"`
	Phone       string     `json:"phone"`
	ParentPhone string     `json:"parent_phone"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
