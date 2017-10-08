package main

import (
	"database/sql"
	"fmt"
	"strconv"
)

func main() {

	database, _ := sql.Open("sqlite3", "./data/mbti_data.db")

	rows, _ := database.Query("SELECT id, mbti_type, posts FROM mbti_data LIMIT 5")
	var id int
	var mbti_type string
	var posts string
	for rows.Next() {
		rows.Scan(&id, &mbti_type, &posts)
		fmt.Println(strconv.Itoa(id) + ": " + mbti_type + " " + posts)
	}

}
