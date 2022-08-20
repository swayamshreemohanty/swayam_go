package album_array_tools

import (
	"errors"
	. "web-server/model"

)

func RemoveAlbumAt(array []AlbumModel, index int) ([]AlbumModel ,*AlbumModel, error) {
	var newArray []AlbumModel
	for arrayIndex,album := range array {
		if arrayIndex == index {
			for i := arrayIndex; i < (len(array) - 1); i++ {
				array[i] = array[i+1]
				newArray = append(newArray, array[i])
			}
			return newArray,&album,nil
		} else {
			newArray = append(newArray, album)
		}
	}
	return array,nil,errors.New("No element found")
}