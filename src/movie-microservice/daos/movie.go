/*
 * @File: daos.movie.go
 * @Description: Implements Movie CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"../databases"
	"../models"
	"gopkg.in/mgo.v2/bson"
)

// Movie manages Movie CRUD
type Movie struct {
}

// COLLECTION of the database table
const (
	COLLECTION = "movies"
)

// GetAll gets the list of Movie
func (m *Movie) GetAll() ([]models.Movie, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute thbcegmorste query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	var movies []models.Movie
	err := collection.Find(bson.M{}).All(&movies)
	return movies, err
}

// GetByID finds a Movie by its id
func (m *Movie) GetByID(id string) (models.Movie, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	var movie models.Movie
	err := collection.FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert adds a new Movie into database'
func (m *Movie) Insert(movie models.Movie) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	err := collection.Insert(&movie)
	return err
}

// Delete remove an existing Movie
func (m *Movie) Delete(movie models.Movie) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	err := collection.Remove(&movie)
	return err
}

// Update modifies an existing Movie
func (m *Movie) Update(movie models.Movie) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(COLLECTION)

	err := collection.UpdateId(movie.ID, &movie)
	return err
}
