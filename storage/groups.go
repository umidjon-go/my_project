package storage

import "edu-crm/models"

func (s *Storage) CreateGroup(group models.Group) (*models.Group, error) {
	var response models.Group
	query := "INSERT INTO groups(course_id, user_id, name, status, days, start_time, end_time, room, start_date, end_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *"
	err := s.DB.QueryRow(query, group.CourseId, group.UserId, group.Name, group.Status, group.Days, group.StartTime, group.EndTime, group.Room, group.StartDate, group.EndDate).
	Scan(&response.Id,
		&response.CourseId,
		&response.UserId,
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
	query := "SELECT * FROM groups WHERE id = $1"
	err := s.DB.QueryRow(query, groupId).
	Scan(&response.Id,
		&response.CourseId,
		&response.UserId,
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
func (s *Storage) PutGroup(groupId int) (*models.Group, error) {
	var response models.Group
	query := "UPDATE groups SET status = 'New' WHERE id = $1 RETURNING *"
	err := s.DB.QueryRow(query, groupId).
	Scan(&response.Id,
		&response.CourseId,
		&response.UserId,
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
func (s *Storage) DeleteGroup(groupId int) (error) {
	query := "DELETE FROM groups WHERE id = $1"
	_, err := s.DB.Exec(query, groupId)
	if err != nil {
		return err
	}
	return nil
}