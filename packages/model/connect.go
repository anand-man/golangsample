package model

import (
	"database/sql"
	"fmt"
	"log"
	// _ "github.com/lib/pq"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/todos")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	con = db
	return db
}
