//attendance.go

package main

import (
	"time"
)

type Attendance struct {
	ClassId   int       `json:"class_id"`
	StudentId int       `json:"student_id"`
	Time      time.Time `json:"time"`
	Status    string    `json:"status"`
}
