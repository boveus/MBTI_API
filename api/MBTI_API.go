package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	sqlite3 "github.com/mattn/go-sqlite3"
)

var database *sql.DB
var query_data []Data

type Data struct {
	MBTI_type string `json:"MBTI_type"`
	Posts     string `json:"Posts"`
}

func init() {
	sql.Register("go-sqlite3", &sqlite3.SQLiteDriver{})
	database, _ = sql.Open("sqlite3", "./data/mbti_data.db")
}

func main() {
	http.HandleFunc("/MBTI", func(w http.ResponseWriter, r *http.Request) {
		mbti_type := r.URL.Query().Get("type")
		if mbti_type != "" {
			query_by_type(mbti_type)
			query_data, _ := json.Marshal(query_data)
			w.Header().Set("Content-Type", "application/json")
			w.Write(query_data)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
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
		query_data = append(query_data, Data{MBTI_type: mbti_type, Posts: posts})
	}
}
