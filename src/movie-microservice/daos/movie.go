/*
 * @File: daos.movie.go
 * @Description: Implements Movie CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"context"

	"github.com/raycad/go-microservices/tree/master/src/movie-microservice/databases"
	models "github.com/raycad/go-microservices/tree/master/src/movie-microservice/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// Get a collection to execute the query against.
	collection := databases.Database.MgDB.Collection(COLLECTION)

	var movies []models.Movie

	// Perform the find operation to get a cursor
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode all documents from the cursor into the movies slice
	if err := cursor.All(context.Background(), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

// GetByID finds a Movie by its id
func (m *Movie) GetByID(id string) (models.Movie, error) {
	// Convert the string ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Movie{}, err
	}

	// Get a collection to execute the query against.
	collection := databases.Database.MgDB.Collection(COLLECTION)

	// Perform the find operation to get a single document
	var movie models.Movie
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&movie)
	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

// Insert adds a new Movie into the database
func (m *Movie) Insert(movie models.Movie) error {
	// Get a collection to execute the query against.
	collection := databases.Database.MgDB.Collection(COLLECTION)

	// Insert a single document
	_, err := collection.InsertOne(context.Background(), movie)
	return err
}

// Delete removes an existing Movie
func (m *Movie) Delete(movie models.Movie) error {
	// Get a collection to execute the query against.
	collection := databases.Database.MgDB.Collection(COLLECTION)

	// Delete a single document based on a filter
	filter := bson.M{"_id": movie.ID} // Assuming you have an ID field in your Movie model
	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}

// Update modifies an existing Movie
func (m *Movie) Update(movie models.Movie) error {
	// Get a collection to execute the query against.
	collection := databases.Database.MgDB.Collection(COLLECTION)

	// Define a filter to identify the document to be updated
	filter := bson.M{"_id": movie.ID} // Assuming you have an ID field in your Movie model

	// Define an update to set the fields of the document
	update := bson.M{"$set": movie}

	// Perform the update operation
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}
