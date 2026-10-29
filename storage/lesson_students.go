package storage

import "edu-crm/models"

func (s *Storage) CreateLessonStudent(lessonStudent models.Lesson_Student) (*models.Lesson_Student, error) {
	var response models.Lesson_Student
	query := "INSERT INTO lesson_students(lesson_id, student_id, status) VALUES($1, $2, $3) RETURNING *"

	err := s.DB.QueryRow(query, lessonStudent.LessonId, lessonStudent.LessonId, lessonStudent.Status).
		Scan(&response.LessonId,
			&response.StudentId,
			&response.Status,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *Storage) GetLessonStudent(lessonStuId int) (*models.Lesson_Student, error) {
	var response models.Lesson_Student
	query := "SELECT * FROM lesson_students WHERE id = $1"
	err := s.DB.QueryRow(query, lessonStuId).
		Scan(&response.LessonId,
			&response.StudentId,
			&response.Status,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
