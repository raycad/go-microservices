/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package daos

import (
	"context"

	"github.com/raycad/go-microservices/tree/master/src/user-microservice/common"
	"github.com/raycad/go-microservices/tree/master/src/user-microservice/databases"
	"github.com/raycad/go-microservices/tree/master/src/user-microservice/models"
	"github.com/raycad/go-microservices/tree/master/src/user-microservice/utils"
	"gopkg.in/mgo.v2/bson"
)

// User manages User CRUD
type User struct {
	utils *utils.Utils
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]models.User, error) {
	collection := databases.Database.MgDB.Collection(common.ColUsers)

	var users []models.User
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &users)
	return users, err
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (models.User, error) {
	var err error
	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return models.User{}, err
	}

	collection := databases.Database.MgDB.Collection(common.ColUsers)

	var user models.User
	err = collection.FindOne(context.TODO(), bson.M{"_id": bson.ObjectIdHex(id)}).Decode(&user)
	return user, err
}

// DeleteByID finds a User by its ID and removes it
func (u *User) DeleteByID(id string) error {
	err := u.utils.ValidateObjectID(id)
	if err != nil {
		return err
	}

	collection := databases.Database.MgDB.Collection(common.ColUsers)

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Login User
func (u *User) Login(name string, password string) (models.User, error) {
	collection := databases.Database.MgDB.Collection(common.ColUsers)

	var user models.User
	err := collection.FindOne(context.TODO(), bson.M{"name": name, "password": password}).Decode(&user)
	return user, err
}

// Insert adds a new User into the database
func (u *User) Insert(user models.User) error {
	collection := databases.Database.MgDB.Collection(common.ColUsers)

	_, err := collection.InsertOne(context.TODO(), &user)
	return err
}

// Delete removes an existing User
func (u *User) Delete(user models.User) error {
	collection := databases.Database.MgDB.Collection(common.ColUsers)

	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": user.ID})
	return err
}

// Update modifies an existing User
func (u *User) Update(user models.User) error {
	collection := databases.Database.MgDB.Collection(common.ColUsers)

	_, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": user.ID}, &user)
	return err
}
