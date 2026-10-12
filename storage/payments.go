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
	query := "SELECT id, course_id, student_id, group_id, user_id, amount, payment_type, status, paid_at, comment, created_at, updated_at FROM payments WHERE id = $1"
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

func (s *Storage) GetPayments() ([]models.Payment, error) {
	query := "SELECT id, course_id, student_id, group_id, user_id, amount, payment_type, status, paid_at, comment, created_at, updated_at FROM payments ORDER BY id DESC"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var p models.Payment
		err := rows.Scan(
			&p.Id,
			&p.CourseId,
			&p.StudentId,
			&p.GroupId,
			&p.UserId,
			&p.Amount,
			&p.PaymentType,
			&p.Status,
			&p.PaidAt,
			&p.Comment,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}
	if payments == nil {
		payments = []models.Payment{}
	}
	return payments, nil
}

func (s *Storage) PutPayment(paymentId int, payment models.Payment) (*models.Payment, error) {
	var response models.Payment
	query := `UPDATE payments SET course_id = $1, student_id = $2, group_id = $3, user_id = $4, amount = $5, payment_type = $6, status = $7, paid_at = $8, comment = $9, updated_at = NOW() WHERE id = $10 RETURNING *`
	err := s.DB.QueryRow(query, payment.CourseId, payment.StudentId, payment.GroupId, payment.UserId, payment.Amount, payment.PaymentType, payment.Status, payment.PaidAt, payment.Comment, paymentId).
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