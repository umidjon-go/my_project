package storage

import "edu-crm/models"

func (s *Storage) CreateGroup(group models.Group) (*models.Group, error) {
	var response models.Group

	query := "INSERT INTO groups(course_id, user_id, branch_id, name, status, days, start_time, end_time, room, start_date, end_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING *"
	err := s.DB.QueryRow(query, group.CourseId, group.UserId, group.BranchId, group.Name, group.Status, group.Days, group.StartTime, group.EndTime, group.Room, group.StartDate, group.EndDate).
		Scan(&response.Id,
			&response.CourseId,
			&response.UserId,
			&response.BranchId,
			&response.Name,
			&response.Status,
			&response.Days,
			&response.StartTime,
			&response.EndTime,
			&response.Room,
			&response.StartDate,
			&response.EndDate,
			&response.CreatedAt,
			&response.UpdatedAt,
		)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) GetGroup(groupId int) (*models.Group, error) {
	var response models.Group
	query := "SELECT id, course_id, user_id, branch_id, name, status, days, start_time, end_time, room, start_date, end_date, created_at, updated_at FROM groups WHERE id = $1"
	err := s.DB.QueryRow(query, groupId).
		Scan(&response.Id,
			&response.CourseId,
			&response.UserId,
			&response.BranchId,
			&response.Name,
			&response.Status,
			&response.Days,
			&response.StartTime,
			&response.EndTime,
			&response.Room,
			&response.StartDate,
			&response.EndDate,
			&response.CreatedAt,
			&response.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) GetGroups() ([]models.Group, error) {
	query := "SELECT id, course_id, user_id, branch_id, name, status, days, start_time, end_time, room, start_date, end_date, created_at, updated_at FROM groups ORDER BY id DESC"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var g models.Group
		err := rows.Scan(
			&g.Id,
			&g.CourseId,
			&g.UserId,
			&g.BranchId,
			&g.Name,
			&g.Status,
			&g.Days,
			&g.StartTime,
			&g.EndTime,
			&g.Room,
			&g.StartDate,
			&g.EndDate,
			&g.CreatedAt,
			&g.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}
	if groups == nil {
		groups = []models.Group{}
	}
	return groups, nil
}

func (s *Storage) PutGroup(groupId int, group models.Group) (*models.Group, error) {
	var response models.Group

	query := "UPDATE groups SET course_id = $1, user_id = $2, branch_id = $3, name = $4, status = $5, days = $6, start_time = $7, end_time = $8, room = $9, start_date = $10, end_date = $11, updated_at = NOW() WHERE id = $12 RETURNING *"
	err := s.DB.QueryRow(query, group.CourseId, group.UserId, group.BranchId, group.Name, group.Status, group.Days, group.StartTime, group.EndTime, group.Room, group.StartDate, group.EndDate, groupId).
		Scan(&response.Id,
			&response.CourseId,
			&response.UserId,
			&response.BranchId,
			&response.Name,
			&response.Status,
			&response.Days,
			&response.StartTime,
			&response.EndTime,
			&response.Room,
			&response.StartDate,
			&response.EndDate,
			&response.CreatedAt,
			&response.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}



func (s *Storage) DeleteGroup(groupId int) error {
	query := "DELETE FROM groups WHERE id = $1"
	_, err := s.DB.Exec(query, groupId)
	if err != nil {
		return err
	}
	return nil
}
