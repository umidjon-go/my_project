package storage

import "edu-crm/models"

func (s *Storage) CreatePayment(payment models.Payment) (*models.Payment, error) {
	var response models.Payment
	query := "INSERT INTO payments(course_id, student_id, group_id, user_id, amount, payment_type, status, paid_at, comment) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"
	err := s.DB.QueryRow(query, payment.CourseId, payment.StudentId, payment.GroupId, payment.UserId, payment.Amount, payment.PaymentType, payment.Status, payment.PaidAt, payment.Comment).
	Scan(&response.Id,
		&response.CourseId,
		&response.StudentId,
		&response.GroupId,
		&response.UserId,
		&response.Amount,
		&response.PaymentType,
		&response.Status,
		&response.PaidAt,
		&response.Comment,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *Storage) GetPayment(paymentId int) (*models.Payment, error) {
	var response models.Payment
	query := "SELECT * FROM payments WHERE id = $1"
	err := s.DB.QueryRow(query, paymentId).
	Scan(&response.Id,
		&response.CourseId,
		&response.StudentId,
		&response.GroupId,
		&response.UserId,
		&response.Amount,
		&response.PaymentType,
		&response.Status,
		&response.PaidAt,
		&response.Comment,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}


func (s *Storage) PutPayment(paymentId int) (*models.Payment, error) {
	var response models.Payment
	query := "UPDATE payments SET amount = 1300000 WHERE id = $1 RETURNING *"
	err := s.DB.QueryRow(query, paymentId).
	Scan(&response.Id,
		&response.CourseId,
		&response.StudentId,
		&response.GroupId,
		&response.UserId,
		&response.Amount,
		&response.PaymentType,
		&response.Status,
		&response.PaidAt,
		&response.Comment,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}


func (s *Storage) DeletePayment(paymentId int) error {
	query := "DELETE FROM payments WHERE id = $1"
	_, err := s.DB.Exec(query, paymentId)
	if err != nil {
		return err
	}
	return nil
}