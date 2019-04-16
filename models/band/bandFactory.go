package band

import (
	"database/sql"
	"music/helper"
)

func newBand(row *sql.Rows) Band {
	var band Band
	err := row.Scan(&band.ID, &band.Name, &band.Nationality)
	helper.CheckError(err)
	return band
}

func newBandSlice(rows *sql.Rows) []Band {
	var bands []Band
	defer rows.Close()
	for rows.Next() {
		band := newBand(rows)
		bands = append(bands, band)
	}
	return bands
}
