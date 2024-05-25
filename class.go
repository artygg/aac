//class.go

package main

import (
	"database/sql"
	"fmt"
	"time"
)

type Class struct {
	Id          int          `json:"id"`
	CourseId    int          `json:"course_id"`
	Room        string       `json:"room"`
	StartTime   time.Time    `json:"time"`
	EndTime     time.Time    `json:"time"`
	Attendances []Attendance `json:"attendances"`
}

func (class *Class) getAttendencesByClass(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT * FROM `attendance` where `ClassId`= '%v'", class.Id)
	rows, err := db.Query(qwery)
	if err != nil {
		return err
	}
	for rows.Next() {
		var attendance Attendance
		err := rows.Scan(&attendance.ClassId, &attendance.StudentId, &attendance.Time, &attendance.Status)
		if err != nil {
			return err
		}
		class.Attendances = append(class.Attendances, attendance)
	}
	return nil
}
