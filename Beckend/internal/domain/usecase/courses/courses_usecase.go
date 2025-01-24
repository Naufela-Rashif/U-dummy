package usecase

import "context"

type CoursesUsecase interface {
	Create(ctx context.Context, request CoursesUsecaseRequest) CourseResponse
	Update(ctx context.Context)
	Delete(ctx context.Context)
	FindAll(ctx context.Context)
	FindById(ctx context.Context)
}
