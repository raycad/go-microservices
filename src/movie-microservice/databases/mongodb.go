/*
 * @File: databases.mongodb.go
 * @Description: Handles MongoDB connections
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package databases

import (
	"context"
	"fmt"
	"time"

	"github.com/raycad/go-microservices/tree/master/src/movie-microservice/common"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDB         *mongo.Database
	Databasename string
}

// Init initializes mongo database
func (db *MongoDB) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Set your MongoDB connection URI with the appropriate values
	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s", common.Config.MgDbUsername, common.Config.MgDbPassword, common.Config.MgAddrs)

	// Set up options to pass to the Connect function
	clientOptions := options.Client().ApplyURI(connectionURI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Error("Failed to connect to MongoDB:", err)

		return err
	}

	// Ensure the client is connected and ping the server
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error("Failed to ping MongoDB:", err)

		return err
	}

	log.Info("Connected to MongoDB!")

	db.MgDB = client.Database(common.Config.MgDbName)

	return nil
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MgDB != nil {
		db.MgDB.Client().Disconnect(context.Background())
	}
}
