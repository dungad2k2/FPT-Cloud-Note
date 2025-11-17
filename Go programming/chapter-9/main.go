package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
func main(){
	db, err := sql.Open("mysql", "root:1@/gogo")
	if err != nil {
		panic(err)
	}
	log.Println(db)
}
