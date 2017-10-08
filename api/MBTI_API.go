package main

import (
	"database/sql"
	"log"
)

var database *sql.DB
var query_data []Data

type Data struct {
	mbti_type string
	posts     string
}

func init() {
	database, _ = sql.Open("sqlite3", "./data/mbti_data.db")
}

func main() {
}

func query_by_type(mbti_type string) {
	query_data = nil
	rows, _ := database.Query("SELECT mbti_type, posts FROM mbti_data WHERE mbti_type=?", mbti_type)
	var posts string
	for rows.Next() {
		err := rows.Scan(&mbti_type, &posts)
		if err != nil {
			log.Fatal(err)
		}

		query_data = append(query_data, Data{mbti_type: mbti_type, posts: posts})
	}
}
