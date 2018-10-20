/*
 * @File: daos.user_dao.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"../databases"
	"../models"
	"../utils"
	"gopkg.in/mgo.v2/bson"
)

// UserDAO manages User CRUD
type UserDAO struct {
}

// COLLECTION of the database table
const (
	COLLECTION = "users"
)

// GetAll gets the list of Users
func (m *UserDAO) GetAll() ([]models.User, error) {
	sessionCopy := databases.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Databasename).C(COLLECTION)

	var users []models.User
	err := collection.Find(bson.M{}).All(&users)
	return users, err
}

// GetByID finds a User by its id
func (m *UserDAO) GetByID(id string) (models.User, error) {
	var err error
	err = utils.ValidateObjectID(id)
	if err != nil {
		return models.User{}, err
	}

	sessionCopy := databases.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Databasename).C(COLLECTION)

	var user models.User
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// DeleteByID finds a User by its id
func (m *UserDAO) DeleteByID(id string) error {
	var err error
	err = utils.ValidateObjectID(id)
	if err != nil {
		return err
	}

	sessionCopy := databases.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Databasename).C(COLLECTION)

	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Login User
func (m *UserDAO) Login(name string, password string) (models.User, error) {
	sessionCopy := databases.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Databasename).C(COLLECTION)

	var user models.User
	err := collection.Find(bson.M{"$and": []bson.M{bson.M{"name": name}, bson.M{"password": password}}}).One(&user)
	return user, err
}

// Insert adds a new User into database'
func (m *UserDAO) Insert(user models.User) error {
	sessionCopy := databases.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Databasename).C(COLLECTION)

	err := collection.Insert(&user)
	return err
}

// Delete remove an existing User
func (m *UserDAO) Delete(user models.User) error {
	sessionCopy := databases.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Databasename).C(COLLECTION)

	err := collection.Remove(&user)
	return err
}

// Update modifies an existing User
func (m *UserDAO) Update(user models.User) error {
	sessionCopy := databases.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Databasename).C(COLLECTION)

	err := collection.UpdateId(user.ID, &user)
	return err
}
