package app

import (
	"database/sql"
	"time"

	"github.com/faridlan/restful-api-market/helper"
)

func NewDB() *sql.DB {
	// user := os.Getenv("USERNAME")
	// pass := os.Getenv("PASSWORD")
	// port := os.Getenv("PORT")
	// host := os.Getenv("HOST")

	db, err := sql.Open("mysql", "doadmin:AVNS_8TEum2I8sYujL-B@tcp(db-mysql-sgp1-26481-do-user-11599617-0.b.db.ondigitalocean.com:25060)/e_market?parseTime=true")
	// db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/e_market?parseTime=true", user, pass, host, port))
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(60)

	return db
}
