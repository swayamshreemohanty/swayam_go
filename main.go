package main

import (
	"github.com/gin-gonic/gin"
	"web-server/controller"
)


func main(){
	router:=gin.Default()
	albums_routes := router.Group("/albums")
	{
		albums_routes.GET("",controller.GetAllAlbumData)
		albums_routes.GET("/:id", controller.GetAlbumById)
		albums_routes.DELETE("deleteAlbum/:id", controller.DeleteAlbumById)
		albums_routes.PUT("editAlbum/:id", controller.PutAlbumById)
		albums_routes.POST("/addAlbum", controller.PostAlbums)
	}
	router.Run("0.0.0.0:8000")
}