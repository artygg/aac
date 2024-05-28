package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Group struct {
	ID string `json:"id"`
}

func getGroupsByCourseID(db *sql.DB, courseID int) ([]Group, error) {
	query := fmt.Sprintf("SELECT GroupID FROM `courses-groups-bridge` WHERE CourseID = %v", courseID)
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

	var groups []Group

	for rows.Next() {
		var group Group
		if err := rows.Scan(&group.ID); err != nil {
			log.Println("Error scanning row:", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return groups, nil
}
