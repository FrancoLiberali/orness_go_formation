package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/FrancoLiberali/orness_go_formation/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// EX3.3 Write a function that returns the list of albums from db
// or returns an error in case of error
func getAlbums() ([]models.Album, error) {
}

// EX3.5 Write a function that takes an id and returns the album with that id from db
// or returns an error in case of error
func getAlbumByID(id uint) (*models.Album, error) {
}

func getAlbumsHandler(c *gin.Context) {
	// EX2.2 add handler that responds the list of albums in json format
	albums, err := getAlbums()
	// EX3.4 return the albums and a Status OK
	// or the error and Internal Server Error in case of error
	// HINT: use c.IndentedJSON
}

func getAlbumByIDHandler(c *gin.Context) {
	// EX2.3 add handler that responds an album by the id present in the url
	// or an error in case the album is not found
	// HINT you will need to transform a string into a int
	albumIDString := c.Param("id")

	albumID, err := strconv.Atoi(albumIDString)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := getAlbumByID(uint(albumID))
	// EX3.6 return the album and a Status OK
	// or Status Not Found in case the album was not found
	// or Status Internal Server Error in case of another error
	// HINT1: gorm will return gorm.ErrRecordNotFound in case a query does not find results
	// HINT2: to compare errors you need to use errors.Is
}

var db *gorm.DB

func main() {
	// EX3.2 connect to the sqlite database "db"
	// and execute auto-migration for Album

	router := gin.Default()
	// EX2.4 add to the router the following routes:
	// /album to get the list of albums
	// /album/id to get an album by its id
	router.GET("/album", getAlbumsHandler)
	router.GET("/album/:id", getAlbumByIDHandler)
	log.Fatal(router.Run("localhost:8080"))
}
