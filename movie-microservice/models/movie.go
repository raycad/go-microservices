/*
 * @File: models.movie.go
 * @Description: Defines Movie information will be returned to the clients
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package models

import "gopkg.in/mgo.v2/bson"

// Movie information
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	URL         string        `bson:"url" json:"url"`
	CoverImage  string        `bson:"coverImage" json:"coverImage"`
	Description string        `bson:"description" json:"description"`
}

// AddMovie information
type AddMovie struct {
	Name        string `json:"name" example:"Movie Name"`
	URL         string `json:"url" example:"Movie URL"`
	CoverImage  string `json:"coverImage" example:"Movie Cover Image"`
	Description string `json:"description" example:"Movie Description"`
}
