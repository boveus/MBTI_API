package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	sqlite3 "github.com/mattn/go-sqlite3"
)

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

	}

	rows, _ := database.Query("SELECT id, mbti_type, posts FROM mbti_data LIMIT 5")
	var id int
	var mbti_type string
	var posts string
	for rows.Next() {
		rows.Scan(&id, &mbti_type, &posts)
		fmt.Println(strconv.Itoa(id) + ": " + mbti_type + " " + posts)
	}

}
