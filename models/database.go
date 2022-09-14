package models

import "time"

//// TEST db
type Test struct {
	Global    Global     `json:"global"`
	Messages  Messages   `json:"messages"`
	Questions []Question `json:"questions"`
}
type Global struct {
	TestID    int      `json:"test_id"`
	Name      string   `json:"name"`
	Company   string   `json:"company"`
	Timeout   int      `json:"timeout"`
	Randomize bool     `json:"randomize"`
	Callback  []string `json:"callback"`
}
type Messages struct {
	Greeting    string `json:"greeting"`
	Information string `json:"information"`
	Thankyou    string `json:"thankyou"`
	Feedback    string `json:"feedback"`
}
type Question struct {
	ID          int      `json:"id"`
	Topic       string   `json:"topic"`
	Timeout     int      `json:"timeout"`
	Question    string   `json:"question"`
	Information string   `json:"information"`
	Multichoice bool     `json:"multichoice"`
	Answers     []string `json:"answers"`
}

//// User's registed information
type UserInformation struct {
	Global Global `json:"global"`
	Users  Users  `json:"candidate"`
}

type Users struct {
	ID        int    `json:"id" omniempty:"id"`
	FirstName string `json:"first_name" omniempty:"first_name"`
	LastName  string `json:"last_name" `
	Contact   string `json:"contact" omniempty:"contact"`
}

//// Returned answers
type ReturnedAnswer struct {
	Global                  Global                  `json:"global"`
	ReturnedUserInformation ReturnedUserInformation `json:"candidate"`
	Stats                   Stats                   `json:"stats"`
	UserAnswers             []UserAnswer            `json"questions"`
}

type ReturnedUserInformation struct {
	TimeStart    time.Time `json:"time_start"`
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname" `
	Contact      string    `json:"contact"`
	SendFeedback bool      `json:"send_feedback"`
	Feedback     string    `json:"feedback"`
}

type Stats struct {
	TimeStart time.Time `json:"time_start" omniempty:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
}

type UserAnswer struct {
	ID          int       `json:"id"`
	Timeout     int       `json:"timeout"`
	Question    string    `json:"question"`
	Multichoice bool      `json:"multichoice"`
	Topic       string    ` json:"topic"`
	Answers     []string  `json:"answers"`
	Clicks      int       `json:"clicks"`
	Histories   []History `json:"history"`
	Results     []Result  `json:"results"`
	Complete    bool      `json:"completed"`
}

type History struct {
	HistoryID int       `json:"id"`
	Pos       int       `json:"pos"`
	Timestamp time.Time `json:" timestamp"`
}

type Result struct {
	Answer   string `json:"answer"`
	Position int    `json:"position"`
	Result   bool   `json:"result"`
}

type Answer struct {
	AnswerId int    `json:"answer_id"`
	Answer   string `json:"answer"`
}

// type Answers struct {
// 	AnswerId  int    `json:"answer_id"`
// 	Name      string `json:"name"`
// 	Timeout   int    `json:"timeout"`
// 	Randomize bool   `json:"randomize"`
// }
// type Candidate struct {
// 	TimeStart string `json:"time_start"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Contact   string `json:"contact"`
// }
