package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	sqlite3 "github.com/mattn/go-sqlite3"
)

var count int

func main() {
	sql.Register("go-sqlite3", &sqlite3.SQLiteDriver{})
	database, _ := sql.Open("sqlite3", "./data/mbti_data.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS mbti_data (id INTEGER PRIMARY KEY, mbti_type TEXT, posts TEXT)")
	statement.Exec()

	f, _ := os.Open("./data/mbti_1.csv")

	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		statement, _ = database.Prepare("INSERT INTO mbti_data (mbti_type, posts) VALUES (?, ?)")
		statement.Exec(record[0], record[1])
		count++
	}
	fmt.Printf("Successfully added %v rows of data to the database.\n", count)
}
