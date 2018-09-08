package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

// Our anonymous db config object
var (
	dbConfig = struct {
		Host,
		User,
		Password,
		Database string
		Port int
	}{
		GetEnvVariable("PG_HOST", "127.0.0.1"),
		GetEnvVariable("PG_USER", "postgres"),
		GetEnvVariable("PG_PASSWORD", "mysecretpassword"),
        "postgres",
		ToInt(GetEnvVariable("PG_PORT", "5432")),
	}

	db *gorm.DB
)

// SetDbConn ...
// Sets a connection to the db
func SetDbConn() {
	connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Database, dbConfig.Password)

	conn, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	conn.LogMode(true)

	db = conn
}

// GetEnvVariable ...
// Tries to get an environment variable. Returns a default value if not found.
func GetEnvVariable(name string, defaultValue string) string {
	value, isSet := os.LookupEnv(name)

	fmt.Print("Print env variable...")
	fmt.Printf("Name: %s \nValue: %s \n\n", name, value)

	if isSet {
		return value
	}

	return defaultValue
}

// ToInt ...
// Just a simple helper function to convert a string to an int
func ToInt(value string) int {
	s, e := strconv.Atoi(value)

	if e != nil {
		return 0
	}

	return s
}
