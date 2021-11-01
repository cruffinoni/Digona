package database

import (
	"errors"
	"fmt"
	"github.com/cruffinoni/Digona/src/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func New() (*Database, error) {
	var (
		err      error
		database Database
	)
	database.db, err = gorm.Open(sqlite.Open("./database/prod.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to open the database: %v", err))
	}
	err = database.db.AutoMigrate(&models.TableUser{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to migrate the schema: %v", err))
	}
	return &database, nil
}
