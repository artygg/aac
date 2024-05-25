//course.go

package main

type Course struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	TeacherID int    `json:"teacherid"`
}
