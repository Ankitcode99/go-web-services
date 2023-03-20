package controller

import (
	"net/http"

	"ankitcode99.com/golang-gin-poc/entity"
	"ankitcode99.com/golang-gin-poc/services"
	"ankitcode99.com/golang-gin-poc/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

/*
**
here we are creating the struct for implementing the interface methods
*
*/
type controller struct {
	services services.VideoServices
}

var validate *validator.Validate

func New(videoServices services.VideoServices) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		services: videoServices,
	}
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.services.Save(video)
	return nil
}

func (c *controller) FindAll() []entity.Video {
	return c.services.FindAll()
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.services.FindAll()

	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
