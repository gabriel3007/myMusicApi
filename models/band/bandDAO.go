package band

import (
	"music/helper"
	"music/models/dao"
)

var table = "bands"
var fields = []string{"name", "nationality"}

func SelectAll() []Band {
	dao.OpenCon()
	rows := dao.SelectRows(table)
	bands := newBandSlice(rows)
	rows.Close()
	dao.CloseCon()
	return bands
}

func Select(id int) Band {
	dao.OpenCon()
	row := dao.SelectRow(table, id)
	band := newBand(row)
	row.Close()
	dao.CloseCon()
	return band
}

func Insert(band Band) {
	dao.OpenCon()
	sql := helper.CreateInsertQuery(table, fields)
	dao.Insert(sql, band.Name, band.Nationality)
	dao.CloseCon()
}

func Update(band Band, id int) {
	dao.OpenCon()
	sql := helper.CreateUpdateQuery(table, fields)
	dao.Update(sql, band.Name, band.Nationality, id)
	dao.CloseCon()
}

func Delete(id int) {
	dao.OpenCon()
	dao.Delete(table, id)
	dao.CloseCon()
}
