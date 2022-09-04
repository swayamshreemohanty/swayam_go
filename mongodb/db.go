package mongodb

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	. "web-server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlbumMongoContext struct {
	ctx context.Context
	mongoclinet *mongo.Client
}
type AlbumMongoService interface{
	InsertAlbumToDB(*AddAlbumModel) (*AlbumModel,error) 
	GetAllAlbumDataFromDB() ([]AlbumModel,error)
	FindAlbumFromDB(id string)(*AlbumModel,error) 
	DeleteAlbumFromDB(id string) error
	UpdateAlbumOnDB(newAddAlbum *AddAlbumModel,id string)error
}
func AlbumMongoServiceInit(ctx context.Context, mongoclinet *mongo.Client) AlbumMongoService {
	return &AlbumMongoContext{
		ctx:ctx,
		mongoclinet: mongoclinet,
	}
}

func (ac *AlbumMongoContext) UpdateAlbumOnDB(newUpdatedAlbum *AddAlbumModel,id string)error  {
	var dbref =ac.mongoclinet.Database("albumDb").Collection("albums")
	filter:=bson.D{primitive.E{Key:"_id",Value:id}}
	update:=bson.M{"$set": newUpdatedAlbum}
	result,err:=dbref.UpdateOne(ac.ctx,filter,update)
	 if err!=nil {
		return err
	 }else if result.MatchedCount !=1 {
		return errors.New("no matched album found for update")
	 }
	return nil
}

func (ac *AlbumMongoContext) DeleteAlbumFromDB(id string)error  {
	var dbref =ac.mongoclinet.Database("albumDb").Collection("albums")
	filter:=bson.D{primitive.E{Key:"_id",Value:id}}

	result,err:=dbref.DeleteOne(ac.ctx,filter)
	 if err!=nil {
		return err
	 }else if result.DeletedCount !=1 {
		return errors.New("no matched album found for delete")
	 }
	return nil
}

func (ac *AlbumMongoContext) FindAlbumFromDB(id string)(*AlbumModel,error)  {
	var dbref =ac.mongoclinet.Database("albumDb").Collection("albums")
	filter:=bson.D{primitive.E{Key:"_id",Value:id}}
	var albumModel AlbumModel
	fmt.Println(filter)
	fmt.Println(albumModel.Artist)
	err:=dbref.FindOne(ac.ctx,filter).Decode(&albumModel)
	 if err!=nil {
		return nil,errors.New("no album found")
	 }
	return &albumModel,nil
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
		return nil,errors.New("no albums found")
	}

	return albumList,nil
}
