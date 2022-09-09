package models

type Users struct {
	ID        int    `json:"id" omniempty:"id"`
	FirstName string `json:"first_name" omniempty:"first_name"`
	LastName  string `json:"last_name" `
	Email     string `json:"email validate:email"`
}

type Answer struct {
	AnswerId int    `json:"answer_id"`
	Answer   string `json:"answer"`
}

type Test struct {
	Global    Global      `json:"global"`
	Messages  Messages    `json:"messages"`
	Questions []Questions `json:"questions"`
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
type Questions struct {
	ID          int      `json:"id"`
	Topic       string   `json:"topic"`
	Timeout     int      `json:"timeout"`
	Question    string   `json:"question"`
	Information string   `json:"information"`
	Multichoice bool     `json:"multichoice"`
	Answers     []string `json:"answers"`
}

type Answers struct {
	TestID    int    `json:"test_id"`
	Name      string `json:"name"`
	Timeout   int    `json:"timeout"`
	Randomize bool   `json:"randomize"`
}
type Candidate struct {
	TimeStart string `json:"time_start"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Contact   string `json:"contact"`
}

type Stats struct {
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
}

type UserAnswer struct {
	Questions []struct {
		ID          int      `json:"id"`
		Timeout     int      `json:"timeout"`
		Question    string   `json:"question"`
		Multichoice bool     `json:"multichoice"`
		Topic       string   `json:"topic"`
		Answers     []string `json:"answers"`
		Clicks      int      `json:"clicks"`
		Completed   bool     `json:"completed"`
	} `json:"questions"`
}

type History struct {
	ID        int   `json:"id"`
	Pos       int   `json:"pos"`
	Timestamp int64 `json:"timestamp"`
}

type Results struct {
	Answer   string `json:"answer"`
	Position int    `json:"position"`
	Result   bool   `json:"result"`
}
