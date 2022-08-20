package controller

import (
	"fmt"
	"net/http"
	. "web-server/db"

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

// func getAlbumById(c *gin.Context)  {
// 	id:=c.Param("id")

// 	//find the id from the album list
// 	for _,album:= range albumList {
// 		if album.Id==id{
// 			c.IndentedJSON(http.StatusOK,album)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound,responseMessage{Message: "Album not found."}) 
// }

// func postAlbums(c *gin.Context)  {
// 	var newAlbum album_model.AddAlbumModel	
// 	if err:= c.ShouldBind(&newAlbum);
// 	err!=nil{
// 		c.IndentedJSON(http.StatusBadRequest,responseMessage{Message: string(err.Error())}) 
// 		return
// 	}else{
// 		var album album_model.AlbumModel;
// 		album.Id=strconv.Itoa(len(albumList)+1)//Assign the next id to the album
// 		album.Title= newAlbum.Title
// 		album.Artist= newAlbum.Artist
// 		album.Price= newAlbum.Price
// 		albumList=append(albumList, album)
// 		c.JSON(http.StatusCreated,gin.H{
// 			"status":"ok",
// 			"data":album,
// 			}) 
// 		return
// 	}
// }