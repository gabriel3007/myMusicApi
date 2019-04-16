package album

import (
	"database/sql"
	"music/helper"
	"music/models/band"
)

func newAlbum(row *sql.Rows) Album {
	var id int
	var name string
	var year int
	var listenedNumValue int
	var bandID int
	err := row.Scan(&id, &name, &year, &listenedNumValue, &bandID)
	helper.CheckError(err)
	band := band.Select(bandID)
	listened := helper.SetBoolValueToAPI(listenedNumValue)
	return Album{id, name, year, band, listened}
}

func newAlbumSlice(rows *sql.Rows) []Album {
	var albums []Album
	defer rows.Close()
	for rows.Next() {
		album := newAlbum(rows)
		albums = append(albums, album)
	}
	return albums
}
