package album

import (
	"music/helper"
	"music/models/dao"
)

var table = "albums"
var fields = []string{"name", "year", "listened", "bandId"}

func SelectAll() []Album {
	dao.OpenCon()
	rows := dao.SelectRows(table)
	albums := newAlbumSlice(rows)
	rows.Close()
	dao.CloseCon()
	return albums
}

func Select(id int) Album {
	dao.OpenCon()
	row := dao.SelectRow(table, id)
	album := newAlbum(row)
	row.Close()
	dao.CloseCon()
	return album
}

func Insert(album Album) {
	dao.OpenCon()
	sql := helper.CreateInsertQuery(table, fields)
	listened := helper.SetBoolValueToDB(album.Listened)
	dao.Insert(sql, album.Name, album.Year, listened, album.Band.ID)
	dao.CloseCon()
}

func Update(album Album, id int) {
	dao.OpenCon()
	sql := helper.CreateUpdateQuery(table, fields)
	listened := helper.SetBoolValueToDB(album.Listened)
	dao.Update(sql, album.Name, album.Year, listened, album.Band.ID, id)
	dao.CloseCon()
}

func Delete(id int) {
	dao.OpenCon()
	dao.Delete(table, id)
	dao.CloseCon()
}
