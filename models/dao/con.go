package dao

import (
	"database/sql"
	"music/helper"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func OpenCon() {
	var err error
	db, err = sql.Open("mysql", "")
	helper.CheckError(err)
}

func CloseCon() {
	db.Close()
}
