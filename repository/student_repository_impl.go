package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

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

func (r *StudentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id uuid.UUID) (*domain.Student, error) {
	query := `SELECT id,name,identity_number,gender,major,class,religion,created_at,updated_at from student WHERE id = $1`
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var student domain.Student
	if rows.Next() {
		err := rows.Scan(&student.ID, &student.Name, &student.IdentityNumber, &student.Gender, &student.Major, &student.Class, &student.Religion, &student.CreatedAt, &student.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return &student, nil
	} else {
		return nil, errors.New("student is not found.")
	}
}
func (r *StudentRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, id uuid.UUID) error {
	query := `DELETE FROM student where id = $1;`
	res, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row < 1 {
		return errors.New("student is not found")
	}
	return nil
}

func (r *StudentRepositoryImpl) UpdateById(ctx context.Context, tx *sql.Tx, id uuid.UUID, student domain.Student) (*domain.Student, error) {
	query := `UPDATE student SET name = $1,identity_number = $2,gender = $3, major = $4, class = $5,religion = $6,updated_at = $7 where id = $8;`
	res, err := tx.ExecContext(ctx, query, student.Name, student.IdentityNumber, student.Gender, student.Major, student.Class, student.Religion, time.Now().UTC(), id)
	if err != nil {
		return nil, err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if row < 1 {
		return nil, errors.New("student is not found.")
	}
	student.ID = id
	student.UpdatedAt = time.Now().UTC()
	return &student, nil

}
