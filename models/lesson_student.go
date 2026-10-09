package models

type Lesson_Student struct {
	Id        int    `json:"id"`
	LessonId  int    `json:"lesson_id"`
	StudentId int    `json:"student_id"`
	Status    string `json:"status"`
}
