package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Course struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Year      int       `json:"year"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	TeacherID int       `json:"teacher_id"`
}

func (course *Course) getCourse(db *sql.DB) error {
	query := fmt.Sprintf("SELECT Name, TeacherID, StartDate, EndDate, Year FROM courses WHERE ID = %d", course.ID)
	return db.QueryRow(query).Scan(&course.Name, &course.TeacherID, &course.StartDate, &course.EndDate, &course.Year)
}

func getCourse(db *sql.DB, id int) (*Course, error) {
	query := fmt.Sprintf("SELECT ID, Name, TeacherID, StartDate, EndDate, Year FROM courses WHERE ID = %d", id)
	row := db.QueryRow(query)

	var course Course
	if err := row.Scan(&course.ID, &course.Name, &course.TeacherID, &course.StartDate, &course.EndDate, &course.Year); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("course not found")
		}
		return nil, err
	}

	return &course, nil
}

func createCourse(db *sql.DB, name string, year int, startDate, endDate time.Time, teacherID int, groups []string) (bool, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Error starting transaction:", err)
		return false, err
	}

	query := `INSERT INTO courses (Name, Year, StartDate, EndDate, TeacherID) VALUES (?, ?, ?, ?, ?)`
	res, err := tx.Exec(query, name, year, startDate, endDate, teacherID)
	if err != nil {
		tx.Rollback()
		log.Println("Error inserting course:", err)
		return false, err
	}

	courseID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println("Error getting last insert ID:", err)
		return false, err
	}

	for _, groupID := range groups {
		query := `INSERT INTO courses_groups_bridge (Courseid, GroupID) VALUES (?, ?)`
		if _, err := tx.Exec(query, courseID, groupID); err != nil {
			tx.Rollback()
			log.Println("Error inserting into courses_groups_bridge:", err)
			return false, err
		}
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction:", err)
		return false, err
	}

	return true, nil
}
