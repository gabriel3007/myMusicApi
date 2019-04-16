package main

import (
	"fmt"
	"log"
	"music/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	setBandsRoutes(router)
	setAlbumsRoutes(router)
	fmt.Println("Server rodando em localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func setBandsRoutes(router *mux.Router) {
	router.HandleFunc("/bandas", api.GetBands).Methods("GET")
	router.HandleFunc("/bandas", api.CreateBand).Methods("POST")
	router.HandleFunc("/bandas/{id}", api.UpdateBand).Methods("PUT")
	router.HandleFunc("/bandas/{id}", api.GetBand).Methods("GET")
	router.HandleFunc("/bandas/{id}", api.DeleteBand).Methods("DELETE")
}

func setAlbumsRoutes(router *mux.Router) {
	router.HandleFunc("/albuns", api.GetAlbums).Methods("GET")
	router.HandleFunc("/albuns", api.CreateAlbum).Methods("POST")
	router.HandleFunc("/albuns{id}", api.UpdateAlbum).Methods("PUT")
	router.HandleFunc("/albuns/{id}", api.GetAlbum).Methods("GET")
	router.HandleFunc("/albuns/{id}", api.DeleteAlbum).Methods("DELETE")
}
