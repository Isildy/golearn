package config

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	DB, err = sqlx.Open("sqlite3", "./golern.db")
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}

}