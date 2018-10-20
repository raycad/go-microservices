/*
 * @File: main.go
 * @Description: Creates HTTP server & API groups of the Movie Service
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package main

import (
	"io"
	"os"

	"./common"
	"./controllers"
	"./databases"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"

	_ "./docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// var router interface{}
var router *gin.Engine

func initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize mongo database
	err = databases.InitMongoDb()
	if err != nil {
		return err
	}

	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	router = gin.Default()

	return nil
}

// @title MovieManagement Service API Document
// @version 1.0
// @description List APIs of MovieManagement Service
// @termsOfService http://swagger.io/terms/

// @host 107.113.53.47:9000
// @BasePath /api/v1
func main() {
	// Initialize server
	if initServer() != nil {
		return
	}

	defer databases.CloseMongoDb()

	c := controllers.NewController()

	// Simple group: v1
	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", c.Login)
		v1.GET("/movies/list", c.ListMovies)

		// APIs need to use token string
		v1.Use(jwt.Auth(common.Config.JwtSecretPassword))
		v1.POST("/movies", c.AddMovie)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(common.Config.Port)
}
