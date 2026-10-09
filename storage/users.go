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
	query := "SELECT * FROM users WHERE id = $1"
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
func (s *Storage) PutUser(userId int) (*models.User, error) {
	var response models.User
	query := "UPDATE users SET first_name = 'Nodir' WHERE id = $1 RETURNING *"
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
func (s *Storage) DeleteUser(userId int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := s.DB.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}	
