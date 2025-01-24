package courses

import (
	"context"
)

type UsecaseCourses interface {
	Create(ctx context.Context, request CoursesCreateRequest) CourseResponse
	Update(ctx context.Context, request CoursesUpdateRequest) CourseResponse
	Delete(ctx context.Context, coursesId int)
	FindAll(ctx context.Context) []CourseResponse
	FindById(ctx context.Context, coursesId int) CourseResponse
}
