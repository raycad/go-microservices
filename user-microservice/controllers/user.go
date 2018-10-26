/*
 * @File: controllers.user.go
 * @Description: Implements User API logic functions
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package controllers

import (
	"net/http"

	"../common"
	"../daos"
	"../models"
	"../utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

// User manages
type User struct {
	utils   utils.Utils
	userDAO daos.User
}

// Authenticate godoc
// @Summary Check user authentication
// @Description Authenticate user
// @Tags admin
// @Security ApiKeyAuth
// @Accept  multipart/form-data
// @Param user formData string true "Username"
// @Param password formData string true "Password"
// @Failure 401 {object} models.Error
// @Failure 500 {object} models.Error
// @Success 200 {object} models.Token
// @Router /admin/auth [post]
func (u *User) Authenticate(ctx *gin.Context) {
	username := ctx.PostForm("user")
	password := ctx.PostForm("password")

	// var user models.User
	var err error
	_, err = u.userDAO.Login(username, password)

	if err == nil {
		var tokenString string
		// Generate token string
		tokenString, err = u.utils.GenerateJWT(username, "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
			log.Debug("[ERROR]: ", err)
			return
		}

		token := models.Token{tokenString}
		// Return token string to the client
		ctx.JSON(http.StatusOK, token)
	} else {
		ctx.JSON(http.StatusUnauthorized, models.Error{common.StatusCodeUnknown, err.Error()})
	}
}

// AddUser godoc
// @Summary Add a new user
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body models.AddUser true "Add user"
// @Failure 500 {object} models.Error
// @Failure 400 {object} models.Error
// @Success 200 {object} models.Message
// @Router /users [post]
func (u *User) AddUser(ctx *gin.Context) {
	var addUser models.AddUser
	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	if err := addUser.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	user := models.User{bson.NewObjectId(), addUser.Name, addUser.Password}
	err := u.userDAO.Insert(user)
	if err == nil {
		ctx.JSON(http.StatusOK, models.Message{"Successfully"})
		log.Debug("Registered a new user = " + user.Name + ", password = " + user.Password)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// ListUsers godoc
// @Summary List all existing users
// @Description List all existing users
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Failure 500 {object} models.Error
// @Success 200 {array} models.User
// @Router /users/list [get]
func (u *User) ListUsers(ctx *gin.Context) {
	var users []models.User
	var err error
	users, err = u.userDAO.GetAll()

	if err == nil {
		ctx.JSON(http.StatusOK, users)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path string true "User ID"
// @Failure 500 {object} models.Error
// @Success 200 {object} models.User
// @Router /users/detail/{id} [get]
func (u *User) GetUserByID(ctx *gin.Context) {
	var user models.User
	var err error
	id := ctx.Params.ByName("id")
	user, err = u.userDAO.GetByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// GetUserByParams godoc
// @Summary Get a user by ID parameter
// @Description Get a user by ID parameter
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id query string true "User ID"
// @Failure 500 {object} models.Error
// @Success 200 {object} models.User
// @Router /users [get]
func (u *User) GetUserByParams(ctx *gin.Context) {
	var user models.User
	var err error
	id := ctx.Request.URL.Query()["id"][0]
	user, err = u.userDAO.GetByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, user)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// DeleteUserByID godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path string true "User ID"
// @Failure 500 {object} models.Error
// @Success 200 {object} models.Message
// @Router /users/{id} [delete]
func (u *User) DeleteUserByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	err := u.userDAO.DeleteByID(id)

	if err == nil {
		ctx.JSON(http.StatusOK, models.Message{"Successfully"})
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user body models.User true "User ID"
// @Failure 500 {object} models.Error
// @Success 200 {object} models.Message
// @Router /users [patch]
func (u *User) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{common.StatusCodeUnknown, err.Error()})
		return
	}

	err := u.userDAO.Update(user)
	if err == nil {
		ctx.JSON(http.StatusOK, models.Message{"Successfully"})
		log.Debug("Registered a new user = " + user.Name + ", password = " + user.Password)
	} else {
		ctx.JSON(http.StatusInternalServerError, models.Error{common.StatusCodeUnknown, err.Error()})
		log.Debug("[ERROR]: ", err)
	}
}
