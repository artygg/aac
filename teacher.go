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

func (teacher *Teacher) getCourses(db *sql.DB) error {
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

func (teacher *Teacher) getTeacher(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT * FROM `teachers` where email LIKE '%v' LIMIT 1", teacher.Email)
	return db.QueryRow(qwery).Scan(&teacher.Id, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Password, &teacher.RegistrationDate)
}

func registerTeacher(db *sql.DB, email, firstName, lastName, password, registrationDate string) (bool, error) {
    query := `
        INSERT INTO teachers (email, firstName, lastName, password, registrationDate)
        VALUES (?, ?, ?, ?, ?)`
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Println("Error preparing query:", err)
        return false, fmt.Errorf("failed to prepare query: %w", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(email, firstName, lastName, password, registrationDate)
    if err != nil {
        log.Println("Error executing query:", err)
        return false, fmt.Errorf("failed to execute query: %w", err)
    }

    return true, nil
}