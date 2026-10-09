package storage

import (
	"edu-crm/models"

	
)

func (s *Storage) CreateStudent(student models.Student) (*models.Student, error) {
	var response models.Student
	query := "INSERT INTO students(course_id, user_Id, full_name, phone, parent_phone) VALUES($1, $2, $3, $4, $5) RETURNING *"
	err := s.DB.QueryRow(query, student.CourseId, student.UserId, student.FullName, student.Phone, student.ParentPhone).
	Scan(&response.Id,
		&response.CourseId,
		&response.UserId,
		&response.GroupId,
		&response.FullName,
		&response.Phone,
		&response.ParentPhone,
		&response.Status,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		return nil, err 
	}
    return &response, nil
}
func (s *Storage) GetStudent(studentId int) (*models.Student, error) {
	var response models.Student
	query := "SELECT * FROM students WHERE id = $1"
	err := s.DB.QueryRow(query, studentId).
	Scan(&response.Id,
        &response.CourseId,
		&response.UserId,
		&response.GroupId,
		&response.FullName,
		&response.Phone,
		&response.ParentPhone,
		&response.Status,
		&response.CreatedAt,
		&response.UpdatedAt,

	)
	if err != nil {
		return nil, err 
	}
    return &response, nil
}
func (s *Storage) PutStudent(studentId int) (*models.Student, error) {
	var response models.Student
	query := "UPDATE students SET phone = '+998887776666' WHERE id = $1 RETURNING *"
	err := s.DB.QueryRow(query, studentId).
	Scan(&response.Id,
        &response.CourseId,
		&response.UserId,
		&response.GroupId,
		&response.FullName,
		&response.Phone,
		&response.ParentPhone,
		&response.Status,
		&response.CreatedAt,
		&response.UpdatedAt,

	)
	if err != nil {
		return nil, err 
	}
    return &response, nil
}
func (s *Storage) DeleteStudent(studentId int) (error) {
	query := "DELETE FROM students WHERE id = $1"
	_, err := s.DB.Exec(query, studentId)
	if err != nil {
		return err 
	}
    return nil
}