//teacher.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Teacher struct {
	Id               int      `json:"id"`
	FirstName        string   `json:"first_name"`
	LastName         string   `json:"last_name"`
	Email            string   `json:"email"`
	Password         string   `json:"password"`
	RegistrationDate string   `json:"registration_date"`
	Courses          []Course `json:"courses"`
}

func (teacher *Teacher) getCourses(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT `id`, `Name`, `TeacherID`, `StartDate`, `EndDate` FROM `courses` WHERE `TeacherID`= '%v'", teacher.Id)
	rows, err := db.Query(qwery)
	if err != nil {
		return err
	}
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.TeacherID, &course.StartDate, &course.EndDate)
		if err != nil {
			return err
		}
		teacher.Courses = append(teacher.Courses, course)
	}
	return nil
}

func (teacher *Teacher) get(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT * FROM `teachers` where email LIKE '%v' LIMIT 1", teacher.Email)
	return db.QueryRow(qwery).Scan(&teacher.Id, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Password, &teacher.RegistrationDate)
}

func (teacher *Teacher) register(db *sql.DB) error {
	query := `
        INSERT INTO teachers (email, firstName, lastName, password, registrationDate)
        VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error preparing query:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(teacher.Email, teacher.FirstName, teacher.LastName, teacher.Password, time.Now())
	if err != nil {
		log.Println("Error executing query:", err)
		return err
	}

	return nil
}
