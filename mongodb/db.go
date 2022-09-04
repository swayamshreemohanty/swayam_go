package mongodb

import (
	"context"
	"errors"
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
	InsertAlbumToDB(*AddAlbumModel) (*AlbumModel,error) 
	GetAllAlbumDataFromDB() ([]AlbumModel,error)
}
func AlbumMongoServiceInit(ctx context.Context, mongoclinet *mongo.Client) AlbumMongoService {
	return &AlbumMongoContext{
		ctx:ctx,
		mongoclinet: mongoclinet,
	}
}

func (ac *AlbumMongoContext)InsertAlbumToDB(addAlbum *AddAlbumModel)(*AlbumModel,error)  {
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
		return nil,err
	}else{
		return &newAlbum,nil
	}
}

func (ac *AlbumMongoContext)GetAllAlbumDataFromDB()([]AlbumModel,error){
	var albumList [] AlbumModel
	var dbref =ac.mongoclinet.Database("albumDb").Collection("albums")
	cursor,err :=	dbref.Find(ac.ctx,bson.D{{}})

	if err !=nil {
		return nil,err
	}

	for cursor.Next(ac.ctx) {
		var albumModel AlbumModel
		err:=cursor.Decode(&albumModel)
		if err!=nil {
			return nil,err
		}
		albumList = append(albumList, albumModel)
	}
	
	if err:=cursor.Err();err!=nil{
		return nil,err
	}
	cursor.Close(ac.ctx)

	if len(albumList)==0{
		return nil,errors.New("No albums found")
	}

	return albumList,nil
}
