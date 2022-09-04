package mongodb

import (
	"context"
	"strconv"
	. "web-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlbumMongoContext struct {
	ctx context.Context
	mongoclinet *mongo.Client
}
type AlbumMongoService interface{
	CreateAlbum(*AddAlbumModel) (AlbumModel,error) 
}
func AlbumMongoServiceInit(ctx context.Context, mongoclinet *mongo.Client) AlbumMongoService {
	return &AlbumMongoContext{
		ctx:ctx,
		mongoclinet: mongoclinet,
	}
}

func (ac *AlbumMongoContext)CreateAlbum(addAlbum *AddAlbumModel)(AlbumModel,error)  {
	var dbref =ac.mongoclinet.Database("albumDb").Collection("albums")

	 totalDocCount,err:= dbref.CountDocuments(ac.ctx,bson.D{})
	 if err!=nil {
		panic(err)
	 }
	var newAlbum AlbumModel	
	// Set the random id:
	newAlbum.Id=strconv.Itoa(int(totalDocCount)+1)
	newAlbum.Title=addAlbum.Title
	newAlbum.Artist=addAlbum.Artist
	newAlbum.Price=addAlbum.Price
	_,err=dbref.InsertOne(ac.ctx,newAlbum)
	if err!=nil {
		return newAlbum,err
	}else{
		return newAlbum,nil
	}
}
