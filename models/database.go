package models

// // TEST db
type Test struct {
	Global    Global      `json:"global"`
	Messages  Messages    `json:"messages"`
	Questions []Questions `json:"questions"`
}
type Global struct {
	TestID    int      `json:"test_id" `
	Name      string   `json:"name" validate:"required"`
	Company   string   `json:"company" validate:"required"`
	Timeout   int      `json:"timeout" `
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
	ID          int      `json:"id" `
	Topic       string   `json:"topic"  validate:"required"`
	Timeout     int      `json:"timeout"`
	Question    string   `json:"question" validate:"required"`
	Information string   `json:"information" validate:"required"`
	Multichoice bool     `json:"multichoice" validate:"required"`
	Answers     []string `json:"answers" validate:"required"`
}

// // Candidate's registed information
type CandidateInformation struct {
	Global    Global    `json:"global"`
	Candidate Candidate `json:"candidate"`
}

type Candidate struct {
	// ID        int    `json:"id" omniempty:"id"`
	TimeStart int64  `json:"time_start"`
	FirstName string `json:"firstname" omniempty:"firstname" validate:"required" `
	LastName  string `json:"lastname" validate:"required" `
	Contact   string `json:"contact" omniempty:"email" validate:"required"`
}

// // Returned answers
type ReturnedAnswer struct {
	Global                       Global                       `json:"global"`
	ReturnedCandidateInformation ReturnedCandidateInformation `json:"candidate"`
	Stats                        Stats                        `json:"stats"`
	Questions                    []ReturnedQuestion           `json:"questions"`
}

type ReturnedCandidateInformation struct {
	TimeStart    int64  `json:"time_start"`
	FirstName    string `json:"firstname" validate:"required"`
	LastName     string `json:"lastname"  validate:"required"`
	Contact      string `json:"contact" validate:"required"`
	SendFeedback bool   `json:"send_feedback" validate:"required"`
	Feedback     string `json:"feedback" validate:"required"`
}

type Stats struct {
	TimeStart int64 `json:"time_start" `
	TimeEnd   int64 `json:"time_end"`
}

type ReturnedQuestion struct {
	ID          int       `json:"id"`
	Timeout     int       `json:"timeout"`
	Question    string    `json:"question" validate:"required"`
	Multichoice bool      `json:"multichoice" validate:"required"`
	Topic       string    `json:"topic" validate:"required"`
	Answers     []string  `json:"answers" validate:"required"`
	Clicks      int       `json:"clicks" validate:"required"`
	Histories   []History `json:"history" validate:"required"`
	Results     []Result  `json:"results" validate:"required"`
	Complete    bool      `json:"completed" validate:"required"`
}

type History struct {
	HistoryID int   `json:"id" validate:"required"`
	Pos       int   `json:"pos" validate:"required"`
	Timestamp int64 `json:"timestamp" validate:"required"`
}

type Result struct {
	Answer   string `json:"answer" validate:"required"`
	Position int    `json:"position" validate:"required"`
	Result   bool   `json:"result" `
}

type Answer struct {
	AnswerId int `json:"answer_id"`

	Answer string `json:"answer"`
}

type Answers struct {
	AnswersId int    `json:"answers_id"`
	Name      string `json:"name"`
	Timeout   int    `json:"timeout"`
	Randomize bool   `json:"randomize"`
}

// func main() {
// 	validate := validator.New()

// 	err := validate.StructCtx(ctx, Test)
// 	if err != nil {

// 	}

// }
