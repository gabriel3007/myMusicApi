package api

import (
	"encoding/json"
	"music/helper"
	"music/models/album"
	"net/http"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	albums := album.SelectAll()
	json.NewEncoder(w).Encode(albums)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := helper.GetIdFromRequest(r)
	album := album.Select(id)
	json.NewEncoder(w).Encode(album)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := helper.GetIdFromRequest(r)
	album.Delete(id)
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typer", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var a album.Album
	json.NewDecoder(r.Body).Decode(&a)
	album.Insert(a)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Typer", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := helper.GetIdFromRequest(r)
	var a album.Album
	json.NewDecoder(r.Body).Decode(&a)
	album.Update(a, id)
}
