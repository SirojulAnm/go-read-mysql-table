package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type employee struct {
	id   int    `json:"id"`
	name string `json:"name"`
	age  int    `json:"age"`
}

func main() {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/company")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	tampil, err := db.Query("SELECT id, name, age FROM employee")

	if err != nil {
		panic(err.Error())
	}

	for tampil.Next() {
		var row employee

		err = tampil.Scan(&row.id, &row.name, &row.age)
		if err != nil {
			panic(err.Error())
		}

		log.Print("\t", row.id, " ", row.name, " ", row.age)
	}

}
