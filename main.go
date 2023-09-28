package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/FrancoLiberali/orness_go_formation/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// EX3.3 Write a function that returns the list of albums from db
// or returns an error in case of error
func getAlbums() ([]models.Album, error) {
	var albums []models.Album
	err := db.Find(&albums).Error

	return albums, err
}

// EX3.5 Write a function that takes an id and returns the album with that id from db
// or returns an error in case of error
func getAlbumByID(id uint) (*models.Album, error) {
	var album models.Album

	err := db.First(&album, id).Error
	if err != nil {
		return nil, err
	}

	return &album, nil
}

func getAlbumsHandler(c *gin.Context) {
	// EX2.2 add handler that responds the list of albums in json format
	albums, err := getAlbums()
	// EX3.4 return the albums and a Status OK
	// or the error and Internal Server Error in case of error
	// HINT: use c.IndentedJSON
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, album)
	}
}

func healthHandler(c *gin.Context) {
	c.Status(http.StatusOK)
}

var db *gorm.DB

func main() {
	// EX3.2 connect to the sqlite database "db"
	// and execute auto-migration for Album
	var err error
	db, err = gorm.Open(sqlite.Open("db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.Album{},
	)

	router := gin.Default()
	// EX2.4 add to the router the following routes:
	// /album to get the list of albums
	// /album/id to get an album by its id
	router.GET("/album", getAlbumsHandler)
	router.GET("/album/:id", getAlbumByIDHandler)
	router.GET("/health", healthHandler)
	log.Fatal(router.Run("localhost:8080"))
}

// EX4.4 Run linting and fix the problems encountered
// install: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
// run: golangci-lint run
