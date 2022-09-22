package validate

import (
	"example.com/m/v2/models"

	"github.com/go4all/validation/types"
)

type Request struct {
	Test models.Test `json:"test"`
	// CandidateInfo models.CandidateInformation `json:"candidate-info"`
	// ReturnAnswer  models.ReturnedAnswer       `json:"return-answer"`
}

func (request Request) Validation() (types.RuleMap, types.MessageMap) {
	TestRules := types.RuleMap{
		"test_id": {"required"},
		"name":    {"required"},

		"firtname": {"required"},
		"lastname": {"required"},
		"contact":  {"required", "email"},
	}
	TestMessages := types.MessageMap{
		"firstname": {
			"required": "Please enter your name",
		},
		"lastname": {
			"required": "Please enter your last name",
		},
		"contact": {
			"required": "Please enter your name",
			"email":    "Not valid email address",
		},
	}
	return TestRules, TestMessages
}

// func (request Request) CandidateInformationValidate() (types.RuleMap, types.MessageMap) {
// 	CandidateRule := types.RuleMap{
// 		"test_id": {"required"},
// 		"name":    {"required"},

// 		"firstname": {"required"},
// 		"lastname":  {"required"},
// 		"contact":   {"required", "email"},
// 	}
// 	CandidateMessages := types.MessageMap{
// 		"name": {
// 			"required": "Please enter your name",
// 		},
// 		"firstname": {"required": "Please enter your name"},
// 		"lastname":  {"required": "Please enter your last name"},
// 		"contact": {
// 			"required": "Please enter your contact",
// 			"email":    "Not valid email address"},
// 	}
// 	return CandidateRule, CandidateMessages
// }

// func (request Request) ReturnedAnswerValidate() (types.RuleMap, types.MessageMap) {
// 	ReturnedAnswerRule := types.RuleMap{
// 		"test_id": {"required"},
// 		"name":    {"required"},

// 		"firstname": {"required"},
// 		"lastname":  {"required"},
// 		"contact":   {"required", "email"},
// 	}
// 	ReturnedAnswerMessages := types.MessageMap{
// 		"name": {
// 			"required": "Please enter your name",
// 		},
// 		"firstname": {"required": "Please enter your name"},
// 		"lastname":  {"required": "Please enter your last name"},
// 		"contact": {
// 			"required": "Please enter your contact",
// 			"email":    "Not valid email address"},
// 	}
// 	return ReturnedAnswerRule, ReturnedAnswerMessages
// }
