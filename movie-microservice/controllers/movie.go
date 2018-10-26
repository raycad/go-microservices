/*
 * @File: controllers.movie.go
 * @Description: Implements Movie API logic functions
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"../common"
	"../daos"
	"../models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

// Movie manages Movie CRUD
type Movie struct {
	movieDAO daos.Movie
}

// Login godoc
// @Summary Log in to the service
// @Description Log in to the service
// @Tags admin
// @Security ApiKeyAuth
// @Accept  multipart/form-data
// @Param user formData string true "Username"
// @Param password formData string true "Password"
// @Failure 401 {object} models.Error
// @Success 200 {object} models.Token
// @Router /login [post]
func (m *Movie) Login(ctx *gin.Context) {
	username := ctx.PostForm("user")
	password := ctx.PostForm("password")

	formData := url.Values{
		"user":     {username},
		"password": {password},
	}

	var authAddr string = common.Config.AuthAddr + "/api/v1/admin/auth"
	resp, err := http.PostForm(authAddr, formData)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
		return
	}

	if resp.StatusCode == http.StatusOK {
		var token models.Token
		json.NewDecoder(resp.Body).Decode(&token)
		ctx.JSON(http.StatusOK, token)
	} else {
		var e models.Error
		json.NewDecoder(resp.Body).Decode(&e)
		ctx.JSON(resp.StatusCode, e)
	}
}

// AddMovie godoc
// @Summary Add a new movie
// @Description Add a new movie
// @Tags movie
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body models.AddMovie true "Add Movie"
// @Failure 401 {object} models.Error
// @Success 200 {object} models.Message
// @Router /movies [post]
func (m *Movie) AddMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.BindJSON(&movie); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	movie.ID = bson.NewObjectId()
	err := m.movieDAO.Insert(movie)
	if err == nil {
		ctx.JSON(http.StatusOK, models.Message{"Successfully"})
	} else {
		ctx.JSON(http.StatusForbidden, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// ListMovies godoc
// @Summary List all existing Movies
// @Description List all existing Movies
// @Tags movie
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 404 {object} models.Error
// @Success 200 {object} models.Movie
// @Router /movies/list [get]
func (m *Movie) ListMovies(ctx *gin.Context) {
	var movies []models.Movie
	var err error
	movies, err = m.movieDAO.GetAll()

	if err == nil {
		ctx.JSON(http.StatusOK, movies)
	} else {
		ctx.JSON(http.StatusNotFound, models.Error{common.StatusCodeUnknown, "Cannot retrieve movie information"})
		log.Debug("[ERROR]: ", err)
	}
}
