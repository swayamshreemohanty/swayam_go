package main

import (
	"net/http"
	"strconv"
	"web-server/model"
	"github.com/gin-gonic/gin"
)

type responseMessage struct{
	Message string `json:"message"`
}

var albumList=[]album_model.AlbumModel{
	{Id:"1",Title: "Bluetrain",Artist: "Swayam",Price: 56.6},
	{Id:"2",Title: "jayveeru",Artist: "Jonny",Price: 30},
	{Id:"3",Title: "Greenbala",Artist: "Dani",Price: 24.8},
}

func getAlbums(c *gin.Context)  {
	c. IndentedJSON(http.StatusOK,albumList)
}

func getAlbumById(c *gin.Context)  {
	id:=c.Param("id")

	//find the id from the album list
	for _,album:= range albumList {
		if album.Id==id{
			c.IndentedJSON(http.StatusOK,album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,responseMessage{Message: "Album not found."}) 
}

func postAlbums(c *gin.Context)  {
	var newAlbum album_model.AddAlbumModel	
	if err:= c.ShouldBind(&newAlbum);
	err!=nil{
		c.IndentedJSON(http.StatusBadRequest,responseMessage{Message: string(err.Error())}) 
		return
	}else{
		var album album_model.AlbumModel;
		album.Id=strconv.Itoa(len(albumList)+1)//Assign the next id to the album
		album.Title= newAlbum.Title
		album.Artist= newAlbum.Artist
		album.Price= newAlbum.Price
		albumList=append(albumList, album)
		c.JSON(http.StatusCreated,gin.H{
			"status":"ok",
			"data":album,
			}) 
		return
	}
}
func main(){
	router:=gin.Default()
	router.GET("/albums",getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/addalbums", postAlbums)
	router.Run("0.0.0.0:8000")
}