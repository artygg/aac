//class.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Class struct {
	Id          int          `json:"id"`
	CourseId    int          `json:"course_id"`
	Room        string       `json:"room"`
	StartTime   time.Time    `json:"starttime"`
	EndTime     time.Time    `json:"endtime"`
	Attendances []Attendance `json:"attendances"`
}

func (class *Class) getAttendences(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT * FROM `attendance` where `ClassId`= '%v'", class.Id)
	rows, err := db.Query(qwery)
	if err != nil {
		return err
	}
	for rows.Next() {
		var attendance Attendance
		err := rows.Scan(&attendance.ClassID, &attendance.StudentID, &attendance.Time, &attendance.Status)
		if err != nil {
			return err
		}
		class.Attendances = append(class.Attendances, attendance)
	}
	return nil
}

func getClassesByCourseID(db *sql.DB, courseID string) ([]Class, error) {
	query := fmt.Sprintf("SELECT Id, CourseID, Room, StartTime, EndTime FROM classes WHERE CourseID = %s", courseID)
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

	var classes []Class

	for rows.Next() {
		var class Class
		if err := rows.Scan(&class.ID, &class.CourseID, &class.Room, &class.StartTime, &class.EndTime); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return classes, nil
}
