package services

import (
	"ankitcode99.com/golang-gin-poc/entity"
)

type VideoServices interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

/***
* structs are used to implement interfaces
 */
type videoServices struct {
	videos []entity.Video
}

func New() VideoServices {
	return &videoServices{
		videos: make([]entity.Video, 0),
	}
}

/***
* we need to specify the struct to implement the interface
 */
func (service *videoServices) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoServices) FindAll() []entity.Video {
	return service.videos
}
