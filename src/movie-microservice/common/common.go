/*
 * @File: common.common.go
 * @Description: Defines common information of the service
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package common

import (
	"encoding/json"
	"os"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

// Configuration stores setting values
type Configuration struct {
	Port                string `json:"port"`
	EnableGinConsoleLog bool   `json:"enableGinConsoleLog"`
	EnableGinFileLog    bool   `json:"enableGinFileLog"`

	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`

	MgAddrs      string `json:"mgAddrs"`
	MgDbName     string `json:"mgDbName"`
	MgDbUsername string `json:"mgDbUsername"`
	MgDbPassword string `json:"mgDbPassword"`

	AuthAddr          string `json:"authAddr"`
	JwtSecretPassword string `json:"jwtSecretPassword"`
	Issuer            string `json:"issuer"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

// Status Text
const (
	ErrNameEmpty      = "Name is empty"
	ErrPasswordEmpty  = "Password is empty"
	ErrNotObjectIDHex = "String is not a valid hex representation of an ObjectId"
)

// Status Code
const (
	StatusCodeUnknown = -1
	StatusCodeOK      = 1000

	StatusMismatch = 10
)

// LoadConfig loads configuration from the config file
func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	// Setting Service Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   Config.LogFilename,
		MaxSize:    Config.LogMaxSize,    // megabytes after which new file is created
		MaxBackups: Config.LogMaxBackups, // number of backups
		MaxAge:     Config.LogMaxAge,     // days
	})
	log.SetLevel(log.DebugLevel)

	// log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
