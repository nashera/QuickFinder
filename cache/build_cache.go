package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func build() {
	database, _ := sql.Open("sqlite3", "./nraboy.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS search_result (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	_, _ = statement.Exec()
	statement, _ = database.Prepare("INSERT INTO search_result (firstname, lastname) VALUES (?, ?)")
	_, _ = statement.Exec("Nic", "Raboy")
	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		_ = rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}