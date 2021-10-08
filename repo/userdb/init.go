package userdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/moh-fajri/learn-garphql-user/util"
)

//UserDB struct sqlx
var UserDB *sqlx.DB

func InitDbCon() {
	user := fmt.Sprint(os.Getenv("POSTGRE_DB_USER"))
	password := fmt.Sprint(os.Getenv("POSTGRE_DB_PASSWORD"))
	host := fmt.Sprint(os.Getenv("POSTGRE_DB_HOST"))
	port := fmt.Sprint(os.Getenv("POSTGRE_DB_PORT"))
	dbName := fmt.Sprint(os.Getenv("POSTGRE_DB_DATABASE"))

	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", host, port, user, password, dbName)

	driverName := os.Getenv("POSTGRE_DB_DRIVER")
	sqldb, err := sql.Open(driverName, urlConnection)
	if err != nil {
		log.Fatalf("⇨ %s Data source 1 %s:%s , Failed : %s \n", driverName, host, port, err.Error())
	}

	db := sqlx.NewDb(sqldb, driverName)
	if err != nil {
		log.Fatalf("⇨ %s Data source %s:%s , Failed : %s \n", driverName, host, port, err.Error())
	}
	fmt.Printf("⇨ %s Data source %s:%s , Successfully connected! \n", driverName, host, port)

	db.SetMaxOpenConns(util.StringToInt(os.Getenv("POSTGRE_DB_MAX_CONNS")))
	db.SetMaxIdleConns(util.StringToInt(os.Getenv("POSTGRE_DB_MAX_IDLE_CONNS")))
	db.SetConnMaxLifetime(5 * time.Minute)

	UserDB = db
}
