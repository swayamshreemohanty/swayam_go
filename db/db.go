package db

import (
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