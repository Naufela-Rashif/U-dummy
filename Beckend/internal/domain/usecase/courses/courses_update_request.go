package courses

type CoursesUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1,max=200"`
}
