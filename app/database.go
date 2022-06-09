package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/faridlan/restful-api-market/helper"
)

func NewDB() *sql.DB {
	user := os.Getenv("USERNAME")
	pass := os.Getenv("PASSWORD")
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	dbname := os.Getenv("DB")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbname))
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
