package storage

import "edu-crm/models"

func (s *Storage) CreateUser(user models.User) (*models.User, error) {
	var response models.User
	query := "INSERT INTO users(first_name, last_name, phone, status, rating_score, role, position) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *"
	err := s.DB.QueryRow(query, user.FirstName, user.LastName, user.Phone, user.Status, user.RatingScore, user.Role, user.Position).
		Scan(&response.Id,
			&response.FirstName,
			&response.LastName,
			&response.Phone,
			&response.Status,
			&response.RatingScore,
			&response.Role,
			&response.Position,
			&response.CreatedAt,
			&response.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) GetUser(userId int) (*models.User, error) {
	var response models.User
	query := "SELECT id, first_name, last_name, phone, status, rating_score, role, position, created_at, updated_at FROM users WHERE id = $1"
	err := s.DB.QueryRow(query, userId).
		Scan(&response.Id,
			&response.FirstName,
			&response.LastName,
			&response.Phone,
			&response.Status,
			&response.RatingScore,
			&response.Role,
			&response.Position,
			&response.CreatedAt,
			&response.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) GetUsers() ([]models.User, error) {
	query := "SELECT id, first_name, last_name, phone, status, rating_score, role, position, created_at, updated_at FROM users ORDER BY id DESC"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.Id,
			&u.FirstName,
			&u.LastName,
			&u.Phone,
			&u.Status,
			&u.RatingScore,
			&u.Role,
			&u.Position,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if users == nil {
		users = []models.User{}
	}
	return users, nil
}

func (s *Storage) PutUser(userId int, user models.User) (*models.User, error) {
	var response models.User
	query := "UPDATE users SET first_name = $1, last_name = $2, phone = $3, status = $4, rating_score = $5, role = $6, position = $7, updated_at = NOW() WHERE id = $8 RETURNING *"
	err := s.DB.QueryRow(query, user.FirstName, user.LastName, user.Phone, user.Status, user.RatingScore, user.Role, user.Position, userId).
		Scan(&response.Id,
			&response.FirstName,
			&response.LastName,
			&response.Phone,
			&response.Status,
			&response.RatingScore,
			&response.Role,
			&response.Position,
			&response.CreatedAt,
			&response.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) DeleteUser(userId int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := s.DB.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}
