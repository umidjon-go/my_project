package models

import "time"

type Lesson struct {
	Id          int        `json:"id"`
	CourseId    int        `json:"course_id"`
	UserId      int        `json:"user_id"`
	GroupId       int      `json:"group_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}
