package models

import "time"

type Group struct {
	Id        int        `json:"id"`
	CourseId  int        `json:"course_id"`
	UserId    int        `json:"user_id"`
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	Days      string     `json:"days"`
	StartTime string     `json:"start_time"`
	EndTime   string     `json:"end_time"`
	Room      string     `json:"room"`
	StartDate string     `json:"start_date"`
	EndDate   string     `json:"end_date"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
