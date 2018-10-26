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

// Main manages main golang application
type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize mongo database
	err = databases.Database.Init()
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

	m.router = gin.Default()

	return nil
}

// @title MovieManagement Service API Document
// @version 1.0
// @description List APIs of MovieManagement Service
// @termsOfService http://swagger.io/terms/

// @host 107.113.53.47:8809
// @BasePath /api/v1
func main() {
	m := Main{}

	// Initialize server
	if m.initServer() != nil {
		return
	}

	defer databases.Database.Close()

	c := controllers.Movie{}

	// Simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		v1.POST("/login", c.Login)
		v1.GET("/movies/list", c.ListMovies)

		// APIs need to use token string
		v1.Use(jwt.Auth(common.Config.JwtSecretPassword))
		v1.POST("/movies", c.AddMovie)
	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	m.router.Run(common.Config.Port)
}
