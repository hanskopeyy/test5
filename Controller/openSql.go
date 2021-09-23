package controller

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Open() (dbase *sql.DB) {
	driver := goDotEnvVariable("DB_DRIVER")
	username := goDotEnvVariable("DB_USER_NAME")
	password := goDotEnvVariable("DB_PASSWORD")
	address := goDotEnvVariable("DB_ADDRESS")
	database := goDotEnvVariable("DB_DATABASE")
	dbase, err := sql.Open(driver, username+":"+password+"@tcp("+address+")"+"/"+database)
	if err != nil {
		panic(err.Error())
	}
	return dbase
}

func OpenGMAdmin() (dbase *sql.DB) {
	driverGM := goDotEnvVariable("DB_DRIVER")
	usernameGM := goDotEnvVariable("DB_USER_NAME_ADMIN")
	passwordGM := goDotEnvVariable("DB_PASSWORD_ADMIN")
	addressGM := goDotEnvVariable("DB_ADDRESS_ADMIN")
	databaseGM := goDotEnvVariable("DB_DATABASE_ADMIN")
	dbase, err := sql.Open(driverGM, usernameGM+":"+passwordGM+"@tcp("+addressGM+")"+"/"+databaseGM)
	if err != nil {
		panic(err.Error())
	}
	return dbase
}
