package dao

import (
	"context"
	"database/sql"
	"music/helper"
)

func SelectRows(table string) *sql.Rows {
	ctx := context.Background()
	helper.CheckError(db.PingContext(ctx))
	sql := "select * from " + table
	rows, err := db.QueryContext(ctx, sql)
	helper.CheckError(err)
	return rows
}

func SelectRow(table string, id int) *sql.Rows {
	ctx := context.Background()
	helper.CheckError(db.PingContext(ctx))
	sql := "select * from " + table + " where id=?"
	row, err := db.QueryContext(ctx, sql, id)
	helper.CheckError(err)
	row.Next()
	return row
}

func Delete(table string, id int) {
	ctx := context.Background()
	helper.CheckError(db.PingContext(ctx))
	sql := "delete from " + table + " where id=?"
	stmt, err := db.Prepare(sql)
	helper.CheckError(err)
	defer stmt.Close()
	stmt.ExecContext(ctx, id)
}

func Insert(sql string, args ...interface{}) {
	ctx := context.Background()
	helper.CheckError(db.PingContext(ctx))
	stmt, err := db.Prepare(sql)
	helper.CheckError(err)
	defer stmt.Close()
	stmt.ExecContext(ctx, args...)
}

func Update(sql string, args ...interface{}) {
	ctx := context.Background()
	helper.CheckError(db.PingContext(ctx))
	stmt, err := db.Prepare(sql)
	helper.CheckError(err)
	defer stmt.Close()
	stmt.ExecContext(ctx, args...)
}
