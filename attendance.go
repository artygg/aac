package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Attendance struct {
	ClassID   int       `json:"class_id"`
	Status    string    `json:"status"`
	StudentID int       `json:"student_id"`
	Time      time.Time `json:"time"`
}

func getAttendance(db *sql.DB, courseID string) ([]Attendance, error) {
	query := fmt.Sprintf(`
		SELECT a.ClassId, a.Status, a.StudentId, a.Time
		FROM attendances a
		INNER JOIN classes c ON a.ClassId = c.Id
		WHERE c.CourseID = %s`, courseID)

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}()

	var attendanceRecords []Attendance

	for rows.Next() {
		var attendance Attendance
		if err := rows.Scan(&attendance.ClassID, &attendance.Status, &attendance.StudentID, &attendance.Time); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		attendanceRecords = append(attendanceRecords, attendance)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return attendanceRecords, nil
}