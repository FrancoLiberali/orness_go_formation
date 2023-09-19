package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/FrancoLiberali/orness_go_formation/models"
	"github.com/gorilla/mux"
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

func getAlbumsHandler(w http.ResponseWriter, r *http.Request) {
	// EX2.2 add handler that responds the list of albums in json format
}

func getAlbumByIDHandler(w http.ResponseWriter, r *http.Request) {
	// EX2.3 add handler that responds an album by the id present in the url
	// or an error in case the album is not found
	// HINT you will need to transform a string into a int
	// EX1.4 Print in the console the album with the id number 1 or the error if produced
	// album, err := getAlbumByID(1)
	// if err != nil {
	// 	log.Println(err.Error())
	// } else {
	// 	log.Println(album)
	// }
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggerMiddleware)
	// EX2.4 add to the router the following routes:
	// /album to get the list of albums
	// /album/id to get an album by its id
	log.Fatal(http.ListenAndServe(":8080", router))
}
