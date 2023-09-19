package models

// EX1.1: Define a Struct Album with the following attributes:
// ID: a numeric identifier of the album
// Name: name of the album
// Artist: artist name
// Songs: amount of songs of the album
// EX2.1: Add json tags to marshall this struct into json
type Album struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Songs  uint   `json:"songs"`
}
