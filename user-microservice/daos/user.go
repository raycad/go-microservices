/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"../common"
	"../databases"
	"../models"
	"../utils"
	"gopkg.in/mgo.v2/bson"
)

// User manages User CRUD
type User struct {
	utils *utils.Utils
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]models.User, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	var users []models.User
	err := collection.Find(bson.M{}).All(&users)
	return users, err
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (models.User, error) {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return models.User{}, err
	}

	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	var user models.User
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return err
	}

	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Login User
func (u *User) Login(name string, password string) (models.User, error) {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	var user models.User
	err := collection.Find(bson.M{"$and": []bson.M{bson.M{"name": name}, bson.M{"password": password}}}).One(&user)
	return user, err
}

// Insert adds a new User into database'
func (u *User) Insert(user models.User) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err := collection.Insert(&user)
	return err
}

// Delete remove an existing User
func (u *User) Delete(user models.User) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err := collection.Remove(&user)
	return err
}

// Update modifies an existing User
func (u *User) Update(user models.User) error {
	sessionCopy := databases.Database.MgDbSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	err := collection.UpdateId(user.ID, &user)
	return err
}
