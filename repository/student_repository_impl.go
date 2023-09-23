package repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/iqbalsonata30/go-student/model/domain"
)

type StudentRepositoryImpl struct{}

func NewRepositoryStudent() StudentRepository {
	return &StudentRepositoryImpl{}
}

func (r *StudentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, student *domain.Student) (*domain.Student, error) {
	query := `INSERT INTO 
    students (name,identity_number,gender,major,class,religion) 
    VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9);
    `
	res, err := tx.ExecContext(ctx, query, student.Name, student.IdentityNumber, student.Gender, student.Major, student.Class, student.Religion)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	student.ID, err = uuid.Parse(strconv.Itoa(int(id)))
	student.CreatedAt = time.Now().UTC()
	student.UpdatedAt = time.Now().UTC()
	if err != nil {
		return nil, err
	}
	return &student, err
}
func (r *StudentRepositoryImpl) FindById(ctx context.Context, id string) (*domain.Student, error) {
	return nil, nil
}
