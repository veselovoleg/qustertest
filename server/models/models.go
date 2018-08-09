package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

//Connection options
var connStr = "host=localhost user=tester password=tester dbname=sitetest sslmode=disable"
var db *sql.DB

//Question struct
type Question struct {
	ID          int    `json:"id"`
	Question    string `json:"question"`
	RightAnswer string `json:"rightanswer"`
}

//QuestionCollection - array of Questions
type QuestionCollection struct {
	Questions []Question `json:"items"`
}

func GetData() QuestionCollection {
	//Connect to DB
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected")
	defer db.Close()

	//Get data from DB
	rows, err := db.Query("SELECT * FROM questions;") //Receive data from content table
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := QuestionCollection{}
	for rows.Next() {
		question := Question{}
		err2 := rows.Scan(&question.ID, &question.Question, &question.RightAnswer)
		if err2 != nil {
			panic(err2)
		}
		result.Questions = append(result.Questions, question)
	}
	return result
}
