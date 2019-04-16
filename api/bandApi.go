package api

import (
	"encoding/json"
	"fmt"
	"music/helper"
	"music/models/band"
	"net/http"
)

func GetBands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	bands := band.SelectAll()
	json.NewEncoder(w).Encode(bands)
}

func GetBand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := helper.GetIdFromRequest(r)
	band := band.Select(id)
	json.NewEncoder(w).Encode(band)
}

func DeleteBand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := helper.GetIdFromRequest(r)
	band.Delete(id)
}

func CreateBand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var b band.Band
	json.NewDecoder(r.Body).Decode(&b)
	fmt.Println("Name:" + b.Name + "|")
	band.Insert(b)
}

func UpdateBand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := helper.GetIdFromRequest(r)
	var b band.Band
	json.NewDecoder(r.Body).Decode(&b)
	band.Update(b, id)
}
