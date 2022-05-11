package app

import (
	"database/sql"
	"time"

	"github.com/faridlan/restful-api-market/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/e_market?parseTime=true")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
