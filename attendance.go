package main

import (
	"database/sql"
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

func getAttendanceByCourse(db *sql.DB, courseID string) ([]Attendance, error) {
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

func updateAttendance(db *sql.DB, studentID, status int) (bool, error) {
	query := `
        UPDATE attendances
        SET Status = ?
        WHERE StudentId = ?`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error preparing query:", err)
		return false, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(status, studentID)
	if err != nil {
		log.Println("Error executing query:", err)
		return false, fmt.Errorf("failed to execute query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return false, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		log.Println("No rows updated")
		return false, fmt.Errorf("no rows updated")
	}

	return true, nil
}
