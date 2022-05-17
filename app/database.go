package app

import (
	"database/sql"
	"time"

	"github.com/faridlan/restful-api-market/helper"
)

func NewDB() *sql.DB {
	// port := os.Getenv("PORT")
	// host := os.Getenv("HOST")

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/e_market?parseTime=true")
	// db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(%s:%s)/belajar_golang", host, port))
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
