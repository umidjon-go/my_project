package storage

import "edu-crm/models"

func (s *Storage) CreateLesson(lesson models.Lesson) (*models.Lesson, error) {
	var response models.Lesson
	query := "INSERT INTO lessons(course_id, user_id, group_id, title, description, status) VALUES($1, $2, $3, $4, $5, COALESCE(NULLIF($6, ''), 'NEW')) RETURNING *"
	err := s.DB.QueryRow(query, lesson.CourseId, lesson.UserId, lesson.GroupId, lesson.Title, lesson.Description, lesson.Status).
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
	query := "SELECT id, course_id, user_id, group_id, title, description, status, created_at, completed_at FROM lessons WHERE id = $1"
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

func (s *Storage) GetLessons() ([]models.Lesson, error) {
	query := "SELECT id, course_id, user_id, group_id, title, description, status, created_at, completed_at FROM lessons ORDER BY id DESC"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []models.Lesson
	for rows.Next() {
		var l models.Lesson
		err := rows.Scan(
			&l.Id,
			&l.CourseId,
			&l.UserId,
			&l.GroupId,
			&l.Title,
			&l.Description,
			&l.Status,
			&l.CreatedAt,
			&l.CompletedAt,
		)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, l)
	}
	if lessons == nil {
		lessons = []models.Lesson{}
	}
	return lessons, nil
}

func (s *Storage) PutLesson(lessonId int, lesson models.Lesson) (*models.Lesson, error) {
	var response models.Lesson
	query := "UPDATE lessons SET course_id = $1, user_id = $2, group_id = $3, title = $4, description = $5, status = $6, updated_at = NOW() WHERE id = $7 RETURNING *"
	err := s.DB.QueryRow(query, lesson.CourseId, lesson.UserId, lesson.GroupId, lesson.Title, lesson.Description, lesson.Status, lessonId).
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

func (s *Storage) CompleteLesson(lessonId int) (*models.Lesson, error) {
	var response models.Lesson
	query := "UPDATE lessons SET status = 'COMPLETED', completed_at = NOW() WHERE id = $1 RETURNING *"
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

func (s *Storage) DeleteLesson(lessonId int) error {
	query := "DELETE FROM lessons WHERE id = $1"
	_, err := s.DB.Exec(query, lessonId)
	if err != nil {
		return err
	}
	return nil
}
