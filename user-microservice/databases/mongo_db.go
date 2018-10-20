/*
 * @File: databases.mongo_db.go
 * @Description: Handles MongoDB connections
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package databases

import (
	"time"

	"../common"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
)

// mgDbSession shares global session
var (
	MgDbSession  *mgo.Session
	Databasename string
)

// InitMongoDb initializes mongo database
func InitMongoDb() error {
	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{common.Config.MgAddrs}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: common.Config.MgDbName,     // Database name
		Username: common.Config.MgDbUsername, // Username
		Password: common.Config.MgDbPassword, // Password
	}

	// Create a session which maintains a pool of socket connections
	// to the DB MongoDB database.
	var err error
	MgDbSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Debug("Can't connect to mongo, go error: ", err)
		return err
	}

	Databasename = common.Config.MgDbName

	// defer MgDbSession.Close()

	return InitData()
}

// InitData initializes default data
func InitData() error {
	var err error

	return err
}

// CloseMongoDb the existing connection
func CloseMongoDb() {
	if MgDbSession != nil {
		MgDbSession.Close()
	}
}
