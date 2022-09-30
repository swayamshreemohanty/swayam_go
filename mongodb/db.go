package mongodb

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	. "web-server/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	UpdateAlbumOnDB(newAddAlbum *AddAlbumModel,id string)(*AlbumModel,error)
}
func AlbumMongoServiceInit(ctx context.Context, mongoclinet *mongo.Client) AlbumMongoService {
	return &AlbumMongoContext{
		ctx:ctx,
		mongoclinet: mongoclinet,
	}
}

func StoreImage(image *multipart.FileHeader) (*string, error)  {
	if(image.Size>10485760){ //10485760 bytes or 10Mb
	return nil,errors.New("file size must be less than 10 Mb")
	}

	src, err:=image.Open()
	if err!=nil {
		return nil,err
	}
	defer src.Close()

	//get current directory path
	mydir, err := os.Getwd()
    if err != nil {
		return nil,err
    }

	
	imageDirectoryPath:= "/album/images/"
	path:=filepath.FromSlash(mydir+imageDirectoryPath)

	//create image path if not present
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil,err
		}
	}
	imageStorePath:= path+filepath.Base(image.Filename)

	dst,err:= os.Create(imageStorePath)
	if err != nil {
		return nil,err
	}
	defer dst.Close()

	_,err= io.Copy(dst,src)	
	if err != nil {
		return nil,err
	}
	imageUrl:="v1/albums"+imageDirectoryPath+filepath.Base(image.Filename)
	return &imageUrl,nil
}

func (albumMongoContext *AlbumMongoContext) UpdateAlbumOnDB(newUpdatedAlbum *AddAlbumModel,id string)(*AlbumModel,error)  {
	var dbref =albumMongoContext.mongoclinet.Database("albumDb").Collection("albums")
	filter:=bson.D{primitive.E{Key:"_id",Value:id}}

	imageUrl,err:=StoreImage(newUpdatedAlbum.Image)
	if err!=nil {
		return nil,err
	}

	var updatedAlbumModel AlbumModel
	updatedAlbumModel.Id=id
	updatedAlbumModel.Title=newUpdatedAlbum.Title
	updatedAlbumModel.Artist=newUpdatedAlbum.Artist
	updatedAlbumModel.Price=newUpdatedAlbum.Price
	updatedAlbumModel.Image=*imageUrl

	update:=bson.M{"$set": updatedAlbumModel}
	result,err:=dbref.UpdateOne(albumMongoContext.ctx,filter,update)
	 if err!=nil {
		return nil,err
	 }else if result.MatchedCount !=1 {
		return nil,errors.New("no matched album found for update")
	 }
	return &updatedAlbumModel,nil
}

func (albumMongoContext *AlbumMongoContext) DeleteAlbumFromDB(id string)error  {
	var dbref =albumMongoContext.mongoclinet.Database("albumDb").Collection("albums")
	filter:=bson.D{primitive.E{Key:"_id",Value:id}}

	result,err:=dbref.DeleteOne(albumMongoContext.ctx,filter)
	 if err!=nil {
		return err
	 }else if result.DeletedCount !=1 {
		return errors.New("no matched album found for delete")
	 }
	return nil
}

func (albumMongoContext *AlbumMongoContext) FindAlbumFromDB(id string)(*AlbumModel,error)  {
	var dbref =albumMongoContext.mongoclinet.Database("albumDb").Collection("albums")
	filter:=bson.D{primitive.E{Key:"_id",Value:id}}
	var albumModel AlbumModel
	err:=dbref.FindOne(albumMongoContext.ctx,filter).Decode(&albumModel)
	 if err!=nil {
		return nil,errors.New("no album found")
	 }
	return &albumModel,nil
}


func (albumMongoContext *AlbumMongoContext) InsertAlbumToDB(addAlbum *AddAlbumModel)(*AlbumModel,error)  {
	var dbref =albumMongoContext.mongoclinet.Database("albumDb").Collection("albums")
	//get the last element of the collection
	var albumModel AlbumModel
	myOption:=options.FindOne()
	myOption.SetSort(bson.M{"$natural":-1})
	dbref.FindOne(albumMongoContext.ctx,bson.M{},myOption).Decode(&albumModel)
	// if err!=nil {
		// return nil,err
	// }
	//
	var newAlbum AlbumModel	
	// increase the id from the last element id:
	lastElementid,err:=strconv.Atoi(albumModel.Id)
	if err!=nil {
		//set the lastElementid to 0, if there is error in find last element, mean the collection is empty
		lastElementid=0
	}
	imageUrl,err:=StoreImage(addAlbum.Image)
	if err!=nil {
		return nil,err
	}
	newAlbum.Image= *imageUrl
	newElementId:=strconv.Itoa(lastElementid+1)
	newAlbum.Id=newElementId
	newAlbum.Title=addAlbum.Title
	newAlbum.Artist=addAlbum.Artist
	newAlbum.Price=addAlbum.Price
	_,err=dbref.InsertOne(albumMongoContext.ctx,newAlbum)
	if err!=nil {
		return nil,err
	}else{
		return &newAlbum,nil
	}
}

func (albumMongoContext *AlbumMongoContext) GetAllAlbumDataFromDB()([]AlbumModel,error){
	var albumList [] AlbumModel
	var dbref =albumMongoContext.mongoclinet.Database("albumDb").Collection("albums")
	cursor,err :=	dbref.Find(albumMongoContext.ctx,bson.D{{}})

	if err !=nil {
		return nil,err
	}

	for cursor.Next(albumMongoContext.ctx) {
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
	cursor.Close(albumMongoContext.ctx)

	if len(albumList)==0{
		return nil,errors.New("no albums found")
	}

	return albumList,nil
}
