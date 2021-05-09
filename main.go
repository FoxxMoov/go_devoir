package main

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"devoir10_ravet/model"
)

// manageErr Manage errors when one is thrown
func manageErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type application struct {
	dataModel *model.DataModel
}

func main() {
	goose.Up("sqlite-database.db")
}

func openDbCon() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite-database.db"), &gorm.Config{})
	manageErr(err)
	db.Exec("PRAGMA foreign_keys = ON")
	db.Debug()
	return db
}
