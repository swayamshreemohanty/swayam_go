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
		// albums_routes.GET("/:id", getAlbumById)
		// albums_routes.POST("", postAlbums)
	}
	router.Run("0.0.0.0:8000")
}