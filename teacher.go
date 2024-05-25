//teacher.go

package main

import (
	"database/sql"
	"fmt"
	"time"
)

type Teacher struct {
	Id               int       `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	RegistrationDate time.Time `json:"registration_date"`
	Courses          []Course  `json:"courses"`
}

func (teacher *Teacher) getCoursesByTeacher(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT * FROM `courses` where `TeacherID`= '%v'", teacher.Id)
	rows, err := db.Query(qwery)
	if err != nil {
		return err
	}
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.Id, &course.Name, &course.TeacherID)
		if err != nil {
			return err
		}
		teacher.Courses = append(teacher.Courses, course)
	}
	return nil
}
