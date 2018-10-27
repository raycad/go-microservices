/*
 * @File: models.user.go
 * @Description: Defines User model
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package models

import (
	"errors"

	"../common"
	"gopkg.in/mgo.v2/bson"
)

// User information
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Name     string        `bson:"name" json:"name" example:"raycad"`
	Password string        `bson:"password" json:"password" example:"raycad"`
}

// AddUser information
type AddUser struct {
	Name     string `json:"name" example:"User Name"`
	Password string `json:"password" example:"User Password"`
}

// Validate user
func (a AddUser) Validate() error {
	switch {
	case len(a.Name) == 0:
		return errors.New(common.ErrNameEmpty)
	case len(a.Password) == 0:
		return errors.New(common.ErrPasswordEmpty)
	default:
		return nil
	}
}
