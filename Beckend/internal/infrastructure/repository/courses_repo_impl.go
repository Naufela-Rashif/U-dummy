package repository

import (
	"Udummy/internal/domain/entity"
	"context"
	"database/sql"
	_ "database/sql/driver"
	"errors"
)

type CoursesRepoImpl struct {
}

func (repo *CoursesRepoImpl) Create(ctx context.Context, tx *sql.Tx, course entity.Courses) entity.Courses {
	SQL := `INSERT INTO courses (id, name,) VALUES (?, ?, ?)`
	result, err := tx.ExecContext(ctx, SQL, course.Name)
	PanicIfError(err)

	id, err := result.LastInsertId()
	PanicIfError(err)

	course.Id = int(id)
	return course
}

func (repo *CoursesRepoImpl) Update(ctx context.Context, tx *sql.Tx, course entity.Courses) entity.Courses {
	SQL := `UPDATE courses SET name = ? WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, course.Name, course.Id)
	PanicIfError(err)
	return course

}

func (repo *CoursesRepoImpl) Delete(ctx context.Context, tx *sql.Tx, course entity.Courses) {
	SQL := `DELETE FROM courses WHERE id = ?`
	_, err := tx.ExecContext(ctx, SQL, course.Id)
	PanicIfError(err)
}

func (repo *CoursesRepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Courses {
	SQL := `SELECT id, name FROM courses`
	rows, err := tx.QueryContext(ctx, SQL)
	PanicIfError(err)

	var courses []entity.Courses
	for rows.Next() {
		course := entity.Courses{}
		err := rows.Scan(&course.Id, &course.Name)
		PanicIfError(err)
		courses = append(courses, course)
	}
	return courses
}

func (repo *CoursesRepoImpl) FindById(ctx context.Context, tx *sql.Tx, courseId int) (entity.Courses, error) {
	SQL := `SELECT id, name FROM courses WHERE id = ?`
	rows, err := tx.QueryContext(ctx, SQL, courseId)
	PanicIfError(err)
	courses := entity.Courses{}
	if rows.Next() {
		err := rows.Scan(&courses.Id, &courses.Name)
		PanicIfError(err)
		return entity.Courses{}, err
	} else {
		return entity.Courses{}, errors.New("courses not found")
	}
}
