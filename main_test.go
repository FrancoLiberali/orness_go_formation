package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/health")
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetAlbums(t *testing.T) {
	// EX4.1 Write a test that gets all the albums and verifies
	// 1. The status code of the response is OK 200
	// 2. The list of albums returned is:
	// 	{
	// 		ID:     1,
	// 		Name:   "Barrio fino",
	// 		Artist: "Daddy yankee",
	// 		Songs:  23,
	// 	},
	// 	{
	// 		ID:     2,
	// 		Name:   "King of Kings",
	// 		Artist: "Don omar",
	// 		Songs:  18,
	// 	},
	// HINT: to verify this last point you will need to unmarshal the json returned by the api
	// you can do it with:
	// 	var albumList []models.Album
	// 	err = json.Unmarshal(body, &albumList)
}

func TestGetAlbumByIDThatExists(t *testing.T) {
	// EX4.2 Write a test that gets the album with id 1 and verifies
	// 1. The status code of the response is OK 200
	// 2. The list of album returned is:
	// 	{
	// 		ID:     1,
	// 		Name:   "Barrio fino",
	// 		Artist: "Daddy yankee",
	// 		Songs:  23,
	// 	}
	// HINT: to verify this last point you will need to unmarshal the json returned by the api
}

func TestGetAlbumByIDThatNotExists(t *testing.T) {
	// EX4.3 Write a test that gets the album with id 3 and verifies
	// 1. The status code of the response is 404 Not Found
}
