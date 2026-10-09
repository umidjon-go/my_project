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
	query := "SELECT * FROM courses WHERE id = $1"
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
func (s *Storage) PutCourse(courseId int) (*models.Course, error) {
	var response models.Course
	query := "UPDATE courses SET price = 1250000  WHERE id = $1 RETURNING *"
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
func (s *Storage) DeleteCourse(courseId int) error {
	query := "DELETE FROM courses where id = $1"
	_, err := s.DB.Exec(query, courseId)
	if err != nil{
		return err
	}
	return nil
}