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
	QuestionID  int      `json:"id"`
	Answer      []Answer `json:"answers"`
	MultiChoice bool     `json:"multichoice"`
	Topic       string   `json:"topic"`
	Information string   `json:"information"`
	Timeouts    int      `json:"timeout"`
	Question    string   `json:"question"`
}

type Test struct {
	TestID    int        `json:"test_id"`
	Name      string     `json:"name"`
	Question  []Question `json:"questions"`
	Randomize bool       `json:"randomize"`
	Company   string     `json:"company"`
	Timeouts  int        `json:"timeout"`
}
