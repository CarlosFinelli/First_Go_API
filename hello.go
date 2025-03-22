package main

import (
	"hello_world/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.GET("/", handler.ReturnHelloWorld)
	router.GET("/", handler.ReturnArtist)
	router.GET("/:id", handler.ReturnArtistById)
	router.POST("/", handler.RegisterAlbum)
	router.PUT("/:id", handler.UpdateAlbum)
	router.DELETE("/:id", handler.DeleteAlbum)

	router.Run("localhost:8080")
}
