package handler

import (
	"fmt"
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
