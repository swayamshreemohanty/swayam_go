package album_model

type AddAlbumModel struct{
	Title string `form:"title" json:"title" binding:"required"`
	Artist string `form:"artist" json:"artist" binding:"required"`
	Price float64 `form:"price" json:"price" binding:"required"`
}


type AlbumModel struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}