package courses

import (
	"Udummy/internal/domain/entity"
)

func ToCoursesResponse(courses entity.Courses) CourseResponse {
	return CourseResponse{
		Id:   courses.Id,
		Name: courses.Name,
	}
}

func ToCourseResponses(courses []entity.Courses) []CourseResponse {
	var coursesResponse []CourseResponse
	for _, course := range courses {
		coursesResponse = append(coursesResponse, ToCoursesResponse(course))
	}
	return coursesResponse
}
