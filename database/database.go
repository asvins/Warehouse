package database

import (
	"github.com/jinzhu/gorm"
	//Blank import needed to call init from pq
	_ "github.com/lib/pq"
)

const (
	databaseType = "postgres"
)

func getUser() string {
	return "postgres"
}

func getDbName() string {
	return "warehouse"
}

func getSSLMode() string {
	return "disable"
}

func buildArgs(user, dbName, sslmode string) string {
	return "user=" + user + " dbname=" + dbName + " sslmode=" + sslmode
}

//GetDatabase ...
func GetDatabase() *gorm.DB {
	user := getUser()
	dbName := getDbName()
	sslmode := getSSLMode()

	db, err := gorm.Open(databaseType, buildArgs(user, dbName, sslmode))
	if err != nil {
		panic(err)
	}
	return &db
}
