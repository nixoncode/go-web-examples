package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "nixon:password@(127.0.0.1:3306)/go_web_eg")
	if err != nil {
		fmt.Printf("Failed to connect to DB: error %s", err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// create a table
	query := `
	CREATE TABLE users(
		id INT AUTO_INCREMENT,
		username VARCHAR(32) NOT NULL,
		password VARCHAR(72) NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print("Table created successfully")
	}
}
