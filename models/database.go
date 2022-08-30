package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Users struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" `
}

type Answer struct {
	AnswerId int64  `json:"answer_id"`
	Answer   string `json:"answer"`
}

type Question struct {
	QuestionID  int64  `json:"question_id"`
	Answer      Answer `json:"answer"`
	MultiChoice bool   `json:"multi_choice"`
	Topic       string `json:"topic"`
	Information string `json:"information"`
	Timeouts    int    `json:"timeouts"`
	Question    string `json:"question"`
	UserId      int    `json:"user_id"`
}

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Users{})

	DB = database
}
