package repository

import (
	"Udummy/internal/domain/entity"
	"context"
	"database/sql"
)

type CoursesRepository interface {
	Create(ctx context.Context, tx *sql.Tx, course entity.Courses) entity.Courses
	Update(ctx context.Context, tx *sql.Tx, course entity.Courses) entity.Courses
	Delete(ctx context.Context, tx *sql.Tx, course entity.Courses)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Courses
	FindById(ctx context.Context, tx *sql.Tx, courseId int) (entity.Courses, error)
}
