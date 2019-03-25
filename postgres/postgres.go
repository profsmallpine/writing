package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/profsmallpine/mid/postgres/migrations"

	// This import is the driver for postgres and required to run a pg db
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConfig used to build database connection.
type DBConfig struct {
	Host       string
	Port       string
	Name       string
	User       string
	Password   string
	VerboseLog bool
}

// Setup runs migrations
func Setup(config *DBConfig) (*gorm.DB, error) {
	db, err := getDBConnection(config)
	if err != nil {
		return nil, err
	}

	if err := migrations.MigrateUp(db); err != nil {
		return db, err
	}

	return db, nil
}

func getDBConnection(config *DBConfig) (*gorm.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Host,
		config.Port,
		config.User,
		config.Name,
		config.Password,
	)

	db, err := gorm.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	if config.VerboseLog {
		db.LogMode(true)
	}
	return db, nil
}
