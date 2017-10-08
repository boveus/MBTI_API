package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

var database *sql.DB

func init() {
	database, _ = sql.Open("sqlite3", "./data/mbti_data.db")
}

func main() {
}

func query_by_type(mbti_type string) {
	rows, _ := database.Query("SELECT id, mbti_type, posts FROM mbti_data LIMIT 5")
	var id int
	var posts string
	for rows.Next() {
		rows.Scan(&id, &mbti_type, &posts)
		fmt.Println(strconv.Itoa(id) + ": " + mbti_type + " " + posts)
	}
}
