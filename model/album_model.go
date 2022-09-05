package album_model

import "mime/multipart"

type AddAlbumModel struct {
	Title  string  `form:"title" binding:"required"`
	Artist string  `form:"artist" binding:"required"`
	Price  float64 `form:"price" binding:"required"`
	Image  *multipart.FileHeader `form:"image" binding:"required"`
}

type AlbumModel struct {
	Id     string  `json:"id" bson:"_id"`
	Title  string  `json:"title" bson:"title"`
	Artist string  `json:"artist" bson:"artist"`
	Price  float64 `json:"price" bson:"price"`
	Image  string 	`json:"image" bson:"image"`
}

type MongoImage struct {
    Id string `bson:"_id"`
}