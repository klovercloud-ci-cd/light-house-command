package config

import (
	"github.com/joho/godotenv"
	"github.com/klovercloud-ci-cd/light-house-command/enums"
	"log"
	"os"
)

// ServerPort refers to server port.
var ServerPort string

// DbServer refers to database server ip.
var DbServer string

// DbPort refers to database server port.
var DbPort string

// DbUsername refers to database name.
var DbUsername string

// DbPassword refers to database password.
var DbPassword string

// DatabaseConnectionString refers to database connection string.
var DatabaseConnectionString string

// DatabaseName refers to database name.
var DatabaseName string

// Database refers to database options.
var Database string

// RunMode refers to run mode.
var RunMode string

// InitEnvironmentVariables initializes environment variables
func InitEnvironmentVariables() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = string(enums.DEVELOP)
	}

	if RunMode != string(enums.PRODUCTION) {
		//Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("ERROR:", err.Error())
			return
		}
	}
	log.Println("RUN MODE:", RunMode)

	ServerPort = os.Getenv("SERVER_PORT")
	DbServer = os.Getenv("MONGO_SERVER")
	DbPort = os.Getenv("MONGO_PORT")
	DbUsername = os.Getenv("MONGO_USERNAME")
	DbPassword = os.Getenv("MONGO_PASSWORD")
	DatabaseName = os.Getenv("DATABASE_NAME")
	Database = os.Getenv("DATABASE")
	if Database == enums.MONGO {
		DatabaseConnectionString = "mongodb://" + DbUsername + ":" + DbPassword + "@" + DbServer + ":" + DbPort
	}
}
