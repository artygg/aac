package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Attendance struct {
	ClassID int     `json:"class_id"`
	Status  string  `json:"status"`
	Student Student `json:"student"`
	Time    string  `json:"time"`
}

func getAttendanceByCourse(db *sql.DB, courseID int) ([]Attendance, error) {
	query := fmt.Sprintf("SELECT a.ClassId, a.Status, a.Time, s.Id, s.FirstName, s.LastName, s.Email FROM attendances a INNER JOIN classes c ON a.ClassId = c.Id INNER JOIN students s ON a.StudentId = s.Id WHERE c.CourseID = %v ", courseID)
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
		if err := rows.Scan(&attendance.ClassID, &attendance.Status, &attendance.Time, &attendance.Student.Id, &attendance.Student.FirstName, &attendance.Student.LastName, &attendance.Student.Email); err != nil {
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

func (attendance *Attendance) update(db *sql.DB) error {
	updateQuery := "UPDATE attendances SET `Status` = ?, `Time` = ? WHERE StudentId = ? AND ClassId = ?"
	stmt, err := db.Prepare(updateQuery)
	if err != nil {
		log.Println("Error preparing update query:", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(attendance.Status, time.Now(), attendance.Student.Id, attendance.ClassID)
	if err != nil {
		log.Println("Error executing update query:", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected by update:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No rows updated. Inserting new row.")

		insertQuery := "INSERT INTO attendances (`Status`, `Time`, `StudentId`, `ClassId`) VALUES (?, ?, ?, ?)"
		insertStmt, err := db.Prepare(insertQuery)
		if err != nil {
			log.Println("Error preparing insert query:", err)
			return err
		}
		defer insertStmt.Close()

		_, err = insertStmt.Exec(attendance.Status, time.Now(), attendance.Student.Id, attendance.ClassID)
		if err != nil {
			log.Println("Error executing insert query:", err)
			return err
		}
	}

	return nil

}
