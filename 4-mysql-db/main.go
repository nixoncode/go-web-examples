package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "nixon:password@(127.0.0.1:3306)/go_web_eg?parseTime=true")
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
	CREATE TABLE IF NOT EXISTS users(
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
		fmt.Println("Table created successfully")
	}

	// insert data
	query = "INSERT INTO users(username, password, created_at) VALUES(?, ?, ?)"

	username := "nikkie"
	password := "unsecure"
	createdAt := time.Now()

	result, err := db.Exec(query, username, password, createdAt)
	if err != nil {
		panic(err)
	}

	userId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Printf("User saved with ID: %d\n", userId)

	// query one user
	var (
		id      int
		user    string
		pass    string
		created time.Time
	)

	query = "SELECT id, username, password, created_at FROM users WHERE id = ?"

	err = db.QueryRow(query, userId).Scan(&id, &user, &pass, &created)
	if err != nil {
		panic(err)
	}

	fmt.Printf("USER INFO: Id:%d Username: %s Password: %s CreatedAt: %s\n", id, user, pass, created)

	// query all users
	query = "SELECT id, username, password, created_at FROM users;"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	type usr struct {
		id        int64
		username  string
		password  string
		createdAt time.Time
	}

	var data []usr
	for rows.Next() {
		var u usr

		err = rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)

		data = append(data, u)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("| ID | Username | Password | CreatedAt |")
	fmt.Println("|----|----------|----------|-----------|")
	for i := range data {
		datum := data[i]
		fmt.Printf("| %d | %s | %s | %s |\n", datum.id, datum.username, datum.password, datum.createdAt)
	}

	// update a row

	query = "UPDATE users SET password = ? WHERE id = ?"

	result, err = db.Exec(query, "plain password", userId)

	if err != nil {
		panic(err)
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated row successfully. Affected rows: %d\n", affectedRows)

	// delete a row
	query = "DELETE FROM users WHERE id = ?"

	result, err = db.Exec(query, userId)
	if err != nil {
		panic(err)
	}

	affectedRows, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Deleted %d row Successfully", affectedRows)
}
