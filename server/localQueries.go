package server

import (
	"log"
	"database/sql"
)

func LocalGetVtubers( db *sql.DB ) ([]Vtuber) {

	rows, err := db.Query("SELECT * FROM vtuber")
	if err != nil {
		log.Fatal(err)
	}
	
	defer rows.Close()

	var vt = []Vtuber {}

	for rows.Next() {
		var chuuba Vtuber
		if err := rows.Scan(&chuuba.ID, &chuuba.Name, &chuuba.Channel, &chuuba.Affiliation); err != nil {
			log.Fatal(err)
		}

		vt = append(vt, chuuba)
	}

	return vt
}