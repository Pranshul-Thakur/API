package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct { //specifies what a field name should be when structs content are put inside the json
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func getalbums(c *gin.Context) { // gets a request and responds with status 200 and a json of all the albums
	c.IndentedJSON(http.StatusOK, albums) // responds with the entire list
}

func postalbums(c *gin.Context) { // gets a request to add a new json entry into the album
	var newalbum album                            // creates empty album struct to store the json data
	if err := c.BindJSON(&newalbum); err != nil { // parses the the json body into newalbum and if already exists, it fails
		return
	}
	albums = append(albums, newalbum)     // adds new album to the album slice
	c.IndentedJSON(http.StatusOK, albums) // responds with the updates list
}

func getalbumsbyid(c *gin.Context) { // gets a request and responds to fetch an album by its id from the json
	id := c.Param("id")        // gets the id parameter from the url /albums/2 -> 2
	for _, a := range albums { // looping through the slice of album
		if a.ID == id { // checks if the current id matches that of the asked one
			c.IndentedJSON(http.StatusOK, a) // if found returns the id as json with status 200
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"}) // else if not found returns the error message album not found
}

var albums = []album{ // data needed in the album slice
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() { // sets up a gin https server on local host and route gets the requests to album to getalbum handler
	router := gin.Default()                  // creates a new gin router with default middleware (logger and recovery)
	router.GET("/albums", getalbums)         // gets request at the getalbums handler
	router.POST("/albums", postalbums)       // gets request at the postalbums handler
	router.GET("/albums/:id", getalbumsbyid) // gets a request at getalbumsbyid handler
	router.Run("localhost:8081")             // starts the server on the given host for incoming requests
}
