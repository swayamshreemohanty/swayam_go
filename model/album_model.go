package album_model

type AddAlbumModel struct {
	Title  string  `form:"title" binding:"required" bson:"title"`
	Artist string  `form:"artist" binding:"required" bson:"artist"`
	Price  float64 `form:"price" binding:"required" bson:"price"`
}

type AlbumModel struct {
	Id     string  `json:"id" bson:"_id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}