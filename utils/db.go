package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

// ConnDB connection db variable
var ConnDB *gorm.DB

// DB connection function
func DB() (db *gorm.DB) {
	var err error
	connect := func() *gorm.DB {

		//open a db connection
		dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "localhost", "postgres", "demo", "postgres") //Build connection string
		// fmt.Println(dbURI)

		db, err = gorm.Open("postgres", dbURI)
		if err != nil {
			log.Println(err)
		}

		return db
	}

	if ConnDB == nil {
		ConnDB = connect()
	} else {
		err = ConnDB.DB().Ping()
	}

	if err != nil {
		ConnDB.Close()
		ConnDB = connect()
	}

	return ConnDB
}
