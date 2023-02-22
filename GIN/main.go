package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums with the list of all albums. as json
func getAlbums(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	//     "status": "success",
	//     "data": albums,
	// })

	c.IndentedJSON(http.StatusOK, albums)
}

// handler func to add more album tracks.
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to newAlbum.
	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}
	fmt.Println(newAlbum)

	// now add the new album to the list of albums.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated , newAlbum)
}

// get album  by ID
func getAlbumsByID(c *gin.Context){
	id := c.Param("id")
	 
	// loop over the list of album toi match with the ID.
	for _ , a := range albums{
		if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
		
	}
	c.IndentedJSON(http.StatusNotFound , gin.H{
		"error": "album not found",
	})
}

func main() {
	fmt.Println("Working with GIN")
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsByID)


	router.Run("localhost:8000")
}
