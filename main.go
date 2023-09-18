package main

import (
	"errors"
	"log"

	"github.com/FrancoLiberali/orness_go_formation/models"
)

// EX1.2 Declare a variable called albums that contains the following reggaeton albums (models.Album):
// 1: 'Barrio fino' by 'Daddy yankee' with 23 songs
// 2: 'King of Kings' by 'Don Omar' with 18 songs
// HINT: use models.Album to create an album, as its definition is in the models package
var albums = []models.Album{
	{
		ID:     1,
		Name:   "Barrio Fino",
		Artist: "Daddy yankee",
		Songs:  23,
	},
	{
		ID:     2,
		Name:   "King of Kings",
		Artist: "Don omar",
		Songs:  18,
	},
}

// EX1.3 Write a function that takes an id and returns the album with that id
// or returns an error in case the album is not found
func getAlbumByID(id uint) (*models.Album, error) {
	for _, album := range albums {
		if album.ID == id {
			return &album, nil
		}
	}

	return nil, errors.New("album not found")
}

func main() {
	// EX1.4 Print in the console the album with the id number 1 or the error if produced
	album, err := getAlbumByID(1)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(album)
	}
}
