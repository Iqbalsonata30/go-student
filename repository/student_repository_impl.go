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
func (r *StudentRepositoryImpl) FindById(ctx context.Context, id string) (*domain.Student, error) {
	return nil, nil
}
