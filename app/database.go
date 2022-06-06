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

	// db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/e_market?parseTime=true")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/e_market?parseTime=true", user, pass, host, port))
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
