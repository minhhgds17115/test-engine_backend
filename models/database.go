package models

type Users struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" `
}

type Answer struct {
	AnswerId int    `json:"answer_id"`
	Answer   string `json:"answer"`
}

type Question struct {
	QuestionID  int      `json:"question_id"`
	Answer      Answer   `json:"answer"`
	MultiChoice bool     `json:"multi_choice"`
	Topic       string   `json:"topic"`
	Information string   `json:"information"`
	Timeouts    int      `json:"timeouts"`
	Question    string   `json:"question"`
	UserId      int      `json:"user_id"`
	History     []string `json:"history"`
}
