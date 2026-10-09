package storage

import (
	"edu-crm/models"

)

func (s Storage) CreateLesson(lesson models.Lesson) (*models.Lesson, error) {
	var response models.Lesson
	query := "INSERT INTO lessons(course_id, user_id, title, description) VALUES($1, $2, $3, $4) RETURNING *"
	err := s.DB.QueryRow(query, lesson.CourseId, lesson.UserId, lesson.Title, lesson.Description).
	Scan(&response.Id,
		&response.CourseId,
		&response.UserId,
		&response.GroupId,
		&response.Title, 
		&response.Description,
		&response.Status,
		&response.CreatedAt,
		&response.CompletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *Storage) GetLesson(lessonId int) (*models.Lesson, error) {
	var response models.Lesson
	query := "SELECT * FROM lessons WHERE id = $1"
	err := s.DB.QueryRow(query, lessonId).
	Scan(&response.Id,
		&response.CourseId,
		&response.UserId,
		&response.GroupId,
		&response.Title, 
		&response.Description,
		&response.Status,
		&response.CreatedAt,
		&response.CompletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *Storage) PutLesson(lessonId int) (*models.Lesson, error) {
	var response models.Lesson
	query := "UPDATE lessons SET id = 1 WHERE id = $1 RETURNING *"
	err := s.DB.QueryRow(query, lessonId).
	Scan(&response.Id,
		&response.CourseId,
		&response.UserId,
		&response.GroupId,
		&response.Title, 
		&response.Description,
		&response.Status,
		&response.CreatedAt,
		&response.CompletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *Storage) DeleteLesson(lessonId int) (error) {
	query := "DELETE FROM lessons WHERE id = $1"
	_, err := s.DB.Exec(query, lessonId)
	if err != nil {
		return err
	}
	return nil
}