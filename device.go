package main

import (
	"database/sql"
	"errors"
)

type Device struct {
	Id string `json:"id"`
	//CurrentLesson Lesson `json:"current_lesson"`
}

func (device *Device) getDevice(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (device *Device) updateDevice(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (device *Device) deleteDevice(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (device *Device) createDevice(db *sql.DB) error {
	return errors.New("Not implemented")
}
