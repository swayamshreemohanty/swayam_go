package album_controller

import (
	"net/http"
	"os"
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

func (albumController *AlbumController) UpdateAlbum(c *gin.Context)  {
	var newAddAlbum AddAlbumModel	
	id:=c.Param("id")

	 err:= c.ShouldBind(&newAddAlbum);
	 if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"false",
			"data":err.Error(),
			}) 
		return
	}else{
		updatedModel,err:=albumController.albumMongoService.UpdateAlbumOnDB(&newAddAlbum,id)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"false",
			"data":err.Error(),
			})
	}else{
		c.JSON(http.StatusAccepted,gin.H{
			"status":"true",
			"data":updatedModel,
		})
		}
	}
}

func (albumController *AlbumController) DeleteAlbumById(c *gin.Context)  {
	id:=c.Param("id")

	err:=albumController.albumMongoService.DeleteAlbumFromDB(id)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"false",
			"data":err.Error(),
			})
	}else{
		c.JSON(http.StatusAccepted,gin.H{
			"status":"true",
			"data":"Id "+id+" removed successfully",
			})
	}
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

func (albumController *AlbumController) GetAlbumImageByName(c *gin.Context){
	filename:=c.Param("filename")

	//get current directory path
	mydir, err := os.Getwd()
    if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
    }
	c.File(mydir+"/album/image/"+filename)
}

func (albumController *AlbumController) GetAlbumById(c *gin.Context){
	id:=c.Param("id")

	albums,err:= albumController.albumMongoService.FindAlbumFromDB(id)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"data":albums,
		})
		return
	}
}

func (albumController *AlbumController) CreateAlbumController(c *gin.Context)  {
	var newAddAlbum AddAlbumModel	
	 err1:= c.ShouldBind(&newAddAlbum);
	 err2:= c.ShouldBindUri(&newAddAlbum);
	 if err1!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"false",
			"data":err1.Error(),
			}) 
		return
	}else if err2!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"false",
			"data":err2.Error(),
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
	albumRoute.DELETE("/deleteAlbum/:id", albumController.DeleteAlbumById)
	albumRoute.GET("/:id", albumController.GetAlbumById)
	albumRoute.PUT("/editAlbum/:id", albumController.UpdateAlbum)

	//get the image
	albumRoute.GET("/album/image/:filename",albumController.GetAlbumImageByName)
}