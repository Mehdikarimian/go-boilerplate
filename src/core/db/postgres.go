package db

import (
	"database/sql"
	"fmt"
	"go-boilerplate/src/config"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorp.DbMap

func ConnectPostgresDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to Postgres!")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap, nil
}

func InitPostgresDB() {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.LoadConfig("POSTGRES_HOST"), config.LoadConfig("POSTGRES_PORT"), config.LoadConfig("POSTGRES_USER"), config.LoadConfig("POSTGRES_PASSWORD"), config.LoadConfig("POSTGRES_DB_NAME"))

	fmt.Println(dbinfo)
	var err error
	db, err = ConnectPostgresDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

func GetPostgresDB() *gorp.DbMap {
	return db
}
