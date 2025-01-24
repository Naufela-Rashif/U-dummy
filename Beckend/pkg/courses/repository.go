package courses

type Repository interface {
	CreateCourse()
	UpdateCourse()
	DeleteCourse()
	FindByIdCourses(id int)
	FindAllCourses()
}
