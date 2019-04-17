package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3361)/mysql")
	if err != nil {
		fmt.Println("error:")
		log.Println(err.Error())
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("select Host,Db from db")
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			Host string
			Db   string
		)
		// var
		if err := rows.Scan(&Host, &Db); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id %s name is %s\n", Host, Db)
	}
}
