package storage

import (
	"edu-crm/models"
)

func (s *Storage) CreateBranch(branch models.Branch) (*models.Branch, error) {
	var response models.Branch
	query := "INSERT INTO branches(name, status, location) VALUES($1, $2, $3) RETURNING *"
	err := s.DB.QueryRow(query, branch.Name, branch.Status, branch.Location).
		Scan(&response.Id,
			&response.Name,
			&response.Status,
			&response.Location,
			&response.CreatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) GetBranch(branchId int) (*models.Branch, error) {
	var response models.Branch
	query := "SELECT id, name, status, location, created_at FROM branches WHERE id = $1"
	err := s.DB.QueryRow(query, branchId).
		Scan(&response.Id,
			&response.Name,
			&response.Status,
			&response.Location,
			&response.CreatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *Storage) GetBranches() ([]models.Branch, error) {
	query := "SELECT id, name, status, location, created_at FROM branches ORDER BY id DESC"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var branches []models.Branch
	for rows.Next() {
		var b models.Branch
		err := rows.Scan(&b.Id, &b.Name, &b.Status, &b.Location, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		branches = append(branches, b)

	}
	if branches == nil {
		branches = []models.Branch{}
	}
	return branches, nil
}
func (s *Storage) PutBranch(branchId int, branch models.Branch) (*models.Branch, error) {
	var response models.Branch
	query := "UPDATE branches SET id = $1, name = $2, status = $3, location = $4 WHERE id = $5 RETURNING *"
	err := s.DB.QueryRow(query, branch.Id, branch.Name, branch.Status, branch.Location, branchId).
		Scan(&response.Id,
			&response.Name,
			&response.Status,
			&response.Location,
			&response.CreatedAt,
		)
	if err != nil {
		return nil, err
	}	
	return &response, nil
}
func (s *Storage) DeleteBranch(branchId int) error {
	query := "DELETE FROM branches WHERE id = $1"
	_, err := s.DB.Exec(query, branchId)
	if err != nil {
		return err
	}
	return nil
}
