package db

import (
	"errors"
	"strconv"
	. "web-server/helper"
	. "web-server/model"
)

var albumList=[]AlbumModel{
	{Id:"1",Title: "Bluetrain",Artist: "Swayam",Price: 56.6},
	{Id:"2",Title: "jayveeru",Artist: "Jonny",Price: 30},
	{Id:"3",Title: "Greenbala",Artist: "Dani",Price: 24.8},
}

type AlbumClient struct{}

type AlbumDB interface{
	GetAllAlbums()(*[]AlbumModel, error)
	StoreAlbumToDB(newRequestedAlbum AddAlbumModel)(*AlbumModel, error)
}

func (_ *AlbumClient) GetAlbumsByIdFromDB(id string)(*AlbumModel, error){

	for _,album:= range albumList {
	//find the id from the album list
		if album.Id==id{
			return &album,nil
		}
	}
	return nil,errors.New("No element found")
}



func (_ *AlbumClient) DeleteAlbumsByIdFromDB(id string)(*AlbumModel, error){
	for albumIndex:= range albumList {
		if albumList[albumIndex].Id==id {
			newList,deletedAlbum,err:=RemoveAlbumAt(albumList,albumIndex)
			if err!=nil {
				return nil,errors.New(err.Error())
			}
			albumList=newList
			return deletedAlbum,nil
		}
	}
	
	return nil,errors.New("No element found")
}

func (_ *AlbumClient) GetAllAlbumsFromDB()(*[]AlbumModel, error){
	result:= make([]AlbumModel, len(albumList))
	var index uint
	for _,v := range albumList {
		result[index]=v
		index++
	}
	return &result,nil
}

func (_ *AlbumClient) StoreAlbumToDB(newRequestedAlbum AddAlbumModel)(*AlbumModel, error){
	
	var lastAlbumId= albumList[len(albumList)-1].Id
	 newAlbumid,err :=strconv.Atoi(lastAlbumId)
	if err!=nil {
		panic("Unable to add the album to data base")
	}
	var newAlbum AlbumModel //creating  the AlbumModel instance
	
	newAlbum.Id=strconv.Itoa(newAlbumid+1) //assign the (lastIndex+1) {ex:albumList last index} as the id of the new element
	
	//Assign AddAlbumModel elements to AlbumModel elements
	newAlbum.Title=newRequestedAlbum.Title
	newAlbum.Artist=newRequestedAlbum.Artist
	newAlbum.Price=newRequestedAlbum.Price
	//

	//Insert the new album to the db
	albumList = append(albumList, newAlbum)

	return &newAlbum,nil
}