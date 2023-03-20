package main

import (
	"io"
	"net/http"
	"os"

	"ankitcode99.com/golang-gin-poc/controller"
	"ankitcode99.com/golang-gin-poc/middlewares"
	"ankitcode99.com/golang-gin-poc/services"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoServices   services.VideoServices     = services.New()
	videoController controller.VideoController = controller.New(videoServices)
)

func setUpLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setUpLogOutput()
	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.CustomLogger(),
		middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/video", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video Input is valid",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", func(ctx *gin.Context) {
			videoController.ShowAll(ctx)
		})
	}

	server.Run(":5000")
}
