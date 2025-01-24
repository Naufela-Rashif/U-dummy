package courses

type CoursesCreateRequest struct {
	Name string `validate:"required,min=1,max=200"`
}
