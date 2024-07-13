package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

type Database struct {
	db *gorm.DB
}

type EntityManagerInterface interface {
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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,         // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	// Get a database handle.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	return db
}

func (d Database) Database() *gorm.DB {
	return d.db
}
