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
				//If any index matched, then just insert/append the newly assigned element till the for loop end.
				newArray = append(newArray, array[i])
			}
			//then return to terminate the function
			return newArray,&album,nil
		} else {
			//insert/append the album from the old array to the new array until if condition satisfied.
			newArray = append(newArray, album)
		}
	}
	return array,nil,errors.New("No element found")
}