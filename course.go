package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Course struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TeacherID int    `json:"teacher_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Year      int    `json:"year"`
}

type Group struct {
	ID string `json:"id"`
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

func createCourse(db *sql.DB, name string, year int, startDate, endDate string, teacherID int, groups []string) (bool, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Error beginning transaction:", err)
		return false, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				log.Println("Error committing transaction:", err)
			}
		}
	}()

	query := `
		INSERT INTO courses (Name, Year, StartDate, EndDate, TeacherID)
		VALUES (?, ?, ?, ?, ?)`
	res, err := tx.Exec(query, name, year, startDate, endDate, teacherID)
	if err != nil {
		log.Println("Error inserting course:", err)
		return false, fmt.Errorf("failed to insert course: %w", err)
	}

	courseID, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		return false, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	for _, groupID := range groups {
		groupQuery := `
			INSERT INTO courses_groups_bridge (Courseid, GroupID)
			VALUES (?, ?)`
		_, err = tx.Exec(groupQuery, courseID, groupID)
		if err != nil {
			log.Println("Error inserting into courses_groups_bridge:", err)
			return false, fmt.Errorf("failed to insert into courses_groups_bridge: %w", err)
		}
	}

	return true, nil
}
