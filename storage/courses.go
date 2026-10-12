package storage

import "edu-crm/models"

func (s *Storage) CreateCourse(course models.Course) (*models.Course, error) {
	var response models.Course
	query := "INSERT INTO courses(user_id, title, description, price, status) VALUES($1, $2, $3, $4, $5) RETURNING *"
	err := s.DB.QueryRow(query, course.UserId, course.Title, course.Description, course.Price, course.Status).
		Scan(&response.Id,
			&response.UserId,
			&response.Title,
			&response.Description,
			&response.Price,
			&response.Status,
			&response.CreatedAt,
			&response.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) GetCourse(courseId int) (*models.Course, error) {
	var response models.Course
	query := "SELECT id, user_id, title, description, price, status, created_at, updated_at FROM courses WHERE id = $1"
	err := s.DB.QueryRow(query, courseId).
		Scan(&response.Id,
			&response.UserId,
			&response.Title,
			&response.Description,
			&response.Price,
			&response.Status,
			&response.CreatedAt,
			&response.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) GetCourses() ([]models.Course, error) {
	query := "SELECT id, user_id, title, description, price, status, created_at, updated_at FROM courses ORDER BY id DESC"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var c models.Course
		err := rows.Scan(&c.Id, &c.UserId, &c.Title, &c.Description, &c.Price, &c.Status, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	if courses == nil {
		courses = []models.Course{}
	}
	return courses, nil
}

func (s *Storage) PutCourse(courseId int, course models.Course) (*models.Course, error) {
	var response models.Course
	query := "UPDATE courses SET user_id = $1, title = $2, description = $3, price = $4, status = $5, updated_at = NOW() WHERE id = $6 RETURNING *"
	err := s.DB.QueryRow(query, course.UserId, course.Title, course.Description, course.Price, course.Status, courseId).
		Scan(&response.Id,
			&response.UserId,
			&response.Title,
			&response.Description,
			&response.Price,
			&response.Status,
			&response.CreatedAt,
			&response.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) DeleteCourse(courseId int) error {
	query := "DELETE FROM courses WHERE id = $1"
	_, err := s.DB.Exec(query, courseId)
	if err != nil {
		return err
	}
	return nil
}