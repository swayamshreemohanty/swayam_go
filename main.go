package main

import (
	"context"
	"fmt"
	"log"
	. "web-server/mongodb"
	. "web-server/album_controller"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx         	context.Context
	mongoClient 	*mongo.Client
	albumController AlbumController
	err  			error

)

func init()  {
	ctx= context.TODO() //non-nil empty context interface for the API calls
	
	mongoConnection:=options.Client().ApplyURI("mongodb://0.0.0.0:27017")
	mongoClient,err=mongo.Connect(ctx,mongoConnection)
	if err !=nil {
		log.Fatal("error while connecting with mongo",err)
	}
	err=mongoClient.Ping(ctx,readpref.Primary())
	if err !=nil {
		log.Fatal("error while tring to ping mongo",err)
	}
	fmt.Println("mongo conection established")
	albumService:=AlbumMongoServiceInit(ctx,mongoClient)
	albumController=New(albumService)
}

func main(){
	defer mongoClient.Disconnect(ctx)
	router:=gin.Default()
	basePath := router.Group("/v1")
	{
		albumController.RegisterAlbumRoutes(basePath)
	}
	router.Run("0.0.0.0:8000")
}