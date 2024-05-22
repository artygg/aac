package main

import (
	"database/sql"
	"fmt"
)

type Teacher struct {
	Id               int    `json:"id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	RegistrationDate string `json:"registration_date"`
}

func (teacher *Teacher) getTeacher(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT * FROM `teachers` where email LIKE '%v' LIMIT 1", teacher.Email)
	return db.QueryRow(qwery).Scan(&teacher.Id, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Password, &teacher.RegistrationDate)
}
