package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type Course struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TeacherID int    `json:"teacher_id"`
}

func (course *Course) getCourse(db *sql.DB) error {
	query := fmt.Sprintf("SELECT Name, TeacherID FROM courses WHERE ID = %d", course.ID)
	return db.QueryRow(query).Scan(&course.Name, &course.TeacherID)
}

func getCourse(db *sql.DB, id int) (*Course, error) {
	query := fmt.Sprintf("SELECT ID, Name, TeacherID FROM courses WHERE ID = %d", id)
	row := db.QueryRow(query)

	var course Course
	if err := row.Scan(&course.ID, &course.Name, &course.TeacherID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("course not found")
		}
		return nil, err
	}

	return &course, nil
}
