package album

import "music/models/band"

type Album struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Year     int       `json:"year"`
	Band     band.Band `json:"band"`
	Listened bool      `json:"listened"`
}
