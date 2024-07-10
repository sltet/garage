package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type Database struct {
	db *gorm.DB
}

type DatabaseInterface interface {
	Database() *gorm.DB
}

func NewDatabase() *Database {
	return &Database{
		db: open(),
	}
}

func open() *gorm.DB {

	params := map[string]string{
		"parseTime": "True",
		"charset":   "utf8mb4",
		"loc":       "Local",
	}

	queryParams := []string{}
	for index, param := range params {
		queryParams = append(queryParams, index+"="+param)
	}
	strings.Join(queryParams[:], "&")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		"avnadmin",
		"AVNS_XXC3taUApAqHGirueob",
		"mysql-garage-14-garage-14.g.aivencloud.com",
		"27893",
		"defaultdb",
		strings.Join(queryParams[:], "&"),
	)

	// Get a database handle.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	return db
}

func (d Database) Database() *gorm.DB {
	return d.db
}
