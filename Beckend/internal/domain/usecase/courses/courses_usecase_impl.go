package courses

import (
	"Udummy/internal/domain/entity"
	"Udummy/internal/domain/repository"
	"Udummy/internal/domain/usecase/helper"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type ServiceImpl struct {
	CoursesRepository repository.CoursesRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func (servise *ServiceImpl) Create(ctx context.Context, request CoursesCreateRequest) CourseResponse {
	err := servise.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := servise.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	course := entity.Courses{
		Name: request.Name,
	}
	course = servise.CoursesRepository.Create(ctx, tx, course)
	return ToCoursesResponse(course)
}

func (servise *ServiceImpl) Update(ctx context.Context, request CoursesUpdateRequest) CourseResponse {
	err := servise.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := servise.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	course, err := servise.CoursesRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	course.Name = request.Name

	course = servise.CoursesRepository.Update(ctx, tx, course)
	return ToCoursesResponse(course)
}

func (servise *ServiceImpl) Delete(ctx context.Context, coursesId int) {
	tx, err := servise.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	course, err := servise.CoursesRepository.FindById(ctx, tx, coursesId)
	helper.PanicIfError(err)

	servise.CoursesRepository.Delete(ctx, tx, course)
}

func (servise *ServiceImpl) FindAll(ctx context.Context) []CourseResponse {
	tx, err := servise.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	courses := servise.CoursesRepository.FindAll(ctx, tx)
	return ToCourseResponses(courses)
}

func (servise *ServiceImpl) FindById(ctx context.Context, coursesId int) CourseResponse {
	tx, err := servise.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	findById, err := servise.CoursesRepository.FindById(ctx, tx, coursesId)
	helper.PanicIfError(err)

	return ToCoursesResponse(findById)
}
