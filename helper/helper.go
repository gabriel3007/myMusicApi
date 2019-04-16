package helper

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Ocorreu o seguinte erro durante a execução: \n\n" + err.Error())
		//panic(err.Error())
	}
}

func CreateInsertQuery(table string, fields []string) string {
	sql := "insert into " + table + " ("
	for i, e := range fields {
		if i < cap(fields)-1 {
			sql += e + ", "
		} else {
			sql += e + ") values ("
		}
	}
	for i := range fields {
		if i < cap(fields)-1 {
			sql += "?,"
		} else {
			sql += "?)"
		}
	}
	return sql
}

func CreateUpdateQuery(table string, fields []string) string {
	sql := "update " + table + " set "
	for i, e := range fields {
		if i < cap(fields)-1 {
			sql += e + "=?, "
		} else {
			sql += e + "=? "
		}
	}
	sql += "where id = ?"
	return sql
}

func SetBoolValueToDB(apiValue bool) int {
	if !apiValue {
		return 0
	}
	return 1
}

func SetBoolValueToAPI(dbValue int) bool {
	if dbValue == 0 {
		return false
	}
	return true
}

func GetIdFromRequest(r *http.Request) int {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	CheckError(err)
	return id
}
