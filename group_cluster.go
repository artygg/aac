//group_cluster.go

package main

import (
	"database/sql"
	"fmt"
)

type GroupCluster struct {
	Groups []Group `json:"groups"`
}

func (groupCluster *GroupCluster) getGroups(db *sql.DB) error {
	qwery := fmt.Sprintf("SELECT * FROM `groups`")
	rows, err := db.Query(qwery)
	if err != nil {
		return err
	}
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.Id)
		if err != nil {
			return err
		}
		groupCluster.Groups = append(groupCluster.Groups, group)
	}
	return nil
}
