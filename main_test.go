package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/FrancoLiberali/orness_go_formation/models"
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
	resp, err := http.Get("http://localhost:8080/album")
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var albumList []models.Album

	err = json.Unmarshal(body, &albumList)
	assert.Nil(t, err)
	assert.Equal(
		t,
		[]models.Album{
			{
				ID:     1,
				Name:   "Barrio fino",
				Artist: "Daddy yankee",
				Songs:  23,
			},
			{
				ID:     2,
				Name:   "King of Kings",
				Artist: "Don omar",
				Songs:  18,
			},
		},
		albumList,
	)
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
	resp, err := http.Get("http://localhost:8080/album/1")
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var album models.Album

	err = json.Unmarshal(body, &album)
	assert.Nil(t, err)
	assert.Equal(
		t,
		models.Album{
			ID:     1,
			Name:   "Barrio fino",
			Artist: "Daddy yankee",
			Songs:  23,
		},
		album,
	)
}

func TestGetAlbumByIDThatNotExists(t *testing.T) {
	// EX4.3 Write a test that gets the album with id 3 and verifies
	// 1. The status code of the response is 404 Not Found
	resp, err := http.Get("http://localhost:8080/album/3")
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
