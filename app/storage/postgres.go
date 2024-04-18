package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tibeahx/growers-dairy/app/common"
)

type DB struct {
	db *sql.DB
}

func NewDB() *DB {
	pgConfig := common.LoadConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pgConfig["host"],
		pgConfig["port"],
		pgConfig["username"],
		pgConfig["password"],
		pgConfig["dbname"])
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		log.Println(err)
	}
	return &DB{
		db: db,
	}
}

// todo: graceful shutdown
func (d *DB) Close() error {
	return d.db.Close()
}
