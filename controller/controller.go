package controller

import (
	"fmt"
	"net/http"
	. "web-server/db"
	. "web-server/model"
	"github.com/gin-gonic/gin"
)

var albumDbClient AlbumClient

func GetAllAlbumData(c *gin.Context){
	albumsList,error:= albumDbClient.GetAllAlbumsFromDB()
	fmt.Print(albumsList)
	if error!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"Error":error.Error()})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"data":albumsList,
		})
		return
	}
}

func DeleteAlbumById(c *gin.Context)  {
	id:=c.Param("id")

	album,err:= albumDbClient.DeleteAlbumsByIdFromDB(id)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"false",
			"data":err.Error(),
			})
	}else{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"true",
			"data":"Id "+album.Id+" removed successfully",
			})
	}
}

func GetAlbumById(c *gin.Context)  {
	id:=c.Param("id")

	album,err:= albumDbClient.GetAlbumsByIdFromDB(id)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"false",
			"data":err.Error(),
			})
	}else{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"true",
			"data":album,
			})
	}
}

func PostAlbums(c *gin.Context)  {
	var newAlbum AddAlbumModel	
	 err:= c.ShouldBind(&newAlbum);
	 if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"false",
			"data":"Unable to add the album",
			}) 
		return
	}else{
		addedAlbum,err:= albumDbClient.StoreAlbumToDB(newAlbum)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":"false",
				"data":err.Error(),
				}) 
			return
		}
		c.JSON(http.StatusCreated,gin.H{
			"status":"ok",
			"data":addedAlbum,
			}) 
		return
	}
}