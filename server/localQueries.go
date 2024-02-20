package server

import (
	"database/sql"
	"log"
)

func LocalGetStreamers(db *sql.DB) []Streamer {

	rows, err := db.Query("SELECT * FROM streamer")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var vt = []Streamer{}

	for rows.Next() {
		var chuuba Streamer
		if err := rows.Scan(&chuuba.ID, &chuuba.Name, &chuuba.Channel, &chuuba.Affiliation); err != nil {
			log.Fatal(err)
		}

		vt = append(vt, chuuba)
	}

	return vt
}
