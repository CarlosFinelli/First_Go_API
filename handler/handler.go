package handler

import (
	"fmt"
	"hello_world/classes"
	"hello_world/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReturnHelloWorld(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.ReturnText())
}

func ReturnArtist(c *gin.Context) {
	param := c.Query("artist")
	name, err := service.ReturnArtist(param)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, name)
}

func ReturnArtistById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	album, erro := service.ReturnById(id)
	if erro != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, album)
}

func RegisterAlbum(c *gin.Context) {
	var album classes.Album

	if err := c.ShouldBindJSON(&album); err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Errorf("error missing parameters: %w", err))
		return
	}

	response, err := service.RegisterAlbum(album)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, response)
}

func UpdateAlbum(c *gin.Context) {
	var album classes.Album
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Errorf("error missing identifier: %w", err))
		return
	}

	if err := c.ShouldBindJSON(&album); err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Errorf("error missing parameters: %w", err))
		return
	}

	response, err := service.UpdateAlbum(id, album)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, fmt.Errorf("internal server error: %v", err))
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

func DeleteAlbum(c *gin.Context) {
	// param := c.Query("id")
	test, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	fmt.Printf("Id: %d", test)
	// id, converterr := strconv.Atoi(param)
	// if converterr != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, fmt.Errorf("unsupported param: %e", converterr))
	// }
	album, err := service.DeleteAlbum(test)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, album)
}
