package models

import "time"

type Courses struct {
	IdCourses     int    `json:"id_courses"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	CourseContent string `json:"course_content"`
	Requirments   string `json:"requirments"`
}

type Users struct {
	ID        int       `json:"id_users"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Password  string    `json:"password"`
}
