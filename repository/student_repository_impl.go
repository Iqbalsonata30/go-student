package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/iqbalsonata30/go-student/model/domain"
)

type StudentRepositoryImpl struct{}

func NewRepositoryStudent() StudentRepository {
	return &StudentRepositoryImpl{}
}

func (r *StudentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, student domain.Student) (*domain.Student, error) {
	query := `INSERT INTO 
    student (id,name,identity_number,gender,major,class,religion) 
    VALUES($1,$2,$3,$4,$5,$6,$7);
    `
	id := uuid.New()
	_, err := tx.ExecContext(ctx, query, id, student.Name, student.IdentityNumber, student.Gender, student.Major, student.Class, student.Religion)
	if err != nil {
		return nil, err
	}
	student.ID = id
	return &student, err
}
func (r *StudentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Student, error) {
	query := `SELECT id,name,identity_number,gender,major,class,religion,created_at,updated_at from student;`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var students []domain.Student
	for rows.Next() {
		var student domain.Student
		if rows.Scan(&student.ID, &student.Name, &student.IdentityNumber, &student.Gender, &student.Major, &student.Class, &student.Religion, &student.CreatedAt, &student.UpdatedAt); err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	defer rows.Close()

	return students, nil

}
