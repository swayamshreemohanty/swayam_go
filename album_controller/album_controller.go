package album_controller

import (
	"net/http"
	. "web-server/model"
	. "web-server/mongodb"
	"github.com/gin-gonic/gin"
)


type AlbumController struct{
	albumMongoService AlbumMongoService
}


func New(albumService AlbumMongoService)  AlbumController{
	return AlbumController{albumMongoService: albumService}
}

func (albumController *AlbumController) GetAllAlbumData(c *gin.Context){
	albumsList,err:= albumController.albumMongoService.GetAllAlbumDataFromDB()
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"data":albumsList,
		})
		return
	}
}

func (albumController *AlbumController)CreateAlbumController(c *gin.Context)  {
	var newAddAlbum AddAlbumModel	
	 err:= c.ShouldBind(&newAddAlbum);
	 if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"false",
			"data":err.Error(),
			}) 
		return
	}else{
		newAddedAlbum,err:= albumController.albumMongoService.InsertAlbumToDB(&newAddAlbum)
		if (err!=nil) {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":"false",
				"data":err.Error(),
				}) 
			return
		}
		c.JSON(http.StatusCreated,gin.H{
			"status":"ok",
			"data":newAddedAlbum,
			}) 
		return
	}
}


func (albumController *AlbumController) RegisterAlbumRoutes(ginRouter *gin.RouterGroup){
	albumRoute:=ginRouter.Group("/albums")
	albumRoute.POST("/create",albumController.CreateAlbumController)
	albumRoute.GET("",albumController.GetAllAlbumData)
	// albums_routes.GET("/:id", controller.GetAlbumById)
	// albums_routes.DELETE("deleteAlbum/:id", controller.DeleteAlbumById)
	// albums_routes.PUT("editAlbum/:id", controller.PutAlbumById)
	// albums_routes.POST("/addAlbum", uc.PostAlbums)
}