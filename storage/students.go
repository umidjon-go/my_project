package storage

import "edu-crm/models"

func (s *Storage) CreateStudent(student models.Student) (*models.Student, error) {
	var response models.Student
	query := "INSERT INTO students(course_id, user_id, group_id, full_name, phone, parent_phone, status) VALUES($1, $2, $3, $4, $5, $6, COALESCE(NULLIF($7, ''), 'NEW')) RETURNING *"
	err := s.DB.QueryRow(query, student.CourseId, student.UserId, student.GroupId, student.FullName, student.Phone, student.ParentPhone, student.Status).
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
	query := "SELECT id, course_id, user_id, group_id, full_name, phone, parent_phone, status, created_at, updated_at FROM students WHERE id = $1"
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

func (s *Storage) GetStudents() ([]models.Student, error) {
	query := "SELECT id, course_id, user_id, group_id, full_name, phone, parent_phone, status, created_at, updated_at FROM students ORDER BY id DESC"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var st models.Student
		err := rows.Scan(
			&st.Id,
			&st.CourseId,
			&st.UserId,
			&st.GroupId,
			&st.FullName,
			&st.Phone,
			&st.ParentPhone,
			&st.Status,
			&st.CreatedAt,
			&st.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, st)
	}
	if students == nil {
		students = []models.Student{}
	}
	return students, nil
}

func (s *Storage) PutStudent(studentId int, student models.Student) (*models.Student, error) {
	var response models.Student
	query := "UPDATE students SET course_id = $1, user_id = $2, group_id = $3, full_name = $4, phone = $5, parent_phone = $6, status = $7, updated_at = NOW() WHERE id = $8 RETURNING *"
	err := s.DB.QueryRow(query, student.CourseId, student.UserId, student.GroupId, student.FullName, student.Phone, student.ParentPhone, student.Status, studentId).
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

func (s *Storage) DeleteStudent(studentId int) error {
	query := "DELETE FROM students WHERE id = $1"
	_, err := s.DB.Exec(query, studentId)
	if err != nil {
		return err
	}
	return nil
}