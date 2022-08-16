package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseMessage struct{
	Message string `json:"message"`
}
type album struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albumList=[]album{
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
	
	var newAlbum album
	if err:= c.BindJSON(&newAlbum);
	err!=nil{
		return
	}
	albumList=append(albumList, newAlbum)
	c.IndentedJSON(http.StatusCreated,responseMessage{Message: "Album added to data base"}) 
}
func main(){
	router:=gin.Default()
	router.GET("/albums",getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/addalbums", postAlbums)
	router.Run("localhost:8080")
}