package request

import (
	"context"

	"github.com/go4all/validation/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestServiceImpl struct {
	requestcollection *mongo.Collection
	ctx               context.Context
}

func NewRequestService(requestcollection *mongo.Collection, ctx context.Context) *RequestServiceImpl {
	return &RequestServiceImpl{
		requestcollection: requestcollection,
		ctx:               ctx,
	}
}
func (r *RequestServiceImpl) ValidationTest() (types.RuleMap, types.MessageMap) {
	TestRules := types.RuleMap{
		"test_id": {"required"},
		"name":    {"required"},
	}
	TestMessages := types.MessageMap{
		"name": {
			"required": "Please enter your name",
		},
	}
	return TestRules, TestMessages
}

func (request *RequestServiceImpl) CandidateInformationValidate() (types.RuleMap, types.MessageMap) {
	CandidateRule := types.RuleMap{
		"test_id": {"required"},
		"name":    {"required"},

		"firstname": {"required"},
		"lastname":  {"required"},
		"contact":   {"required", "email"},
	}
	CandidateMessages := types.MessageMap{
		"name": {
			"required": "Please enter your name",
		},
		"firstname": {"required": "Please enter your name"},
		"lastname":  {"required": "Please enter your last name"},
		"contact": {
			"required": "Please enter your contact",
			"email":    "Not valid email address"},
	}
	return CandidateRule, CandidateMessages
}

func (request *RequestServiceImpl) ReturnedAnswerValidate() (types.RuleMap, types.MessageMap) {
	ReturnedAnswerRule := types.RuleMap{
		"test_id": {"required"},
		"name":    {"required"},

		"firstname": {"required"},
		"lastname":  {"required"},
		"contact":   {"required", "email"},
	}
	ReturnedAnswerMessages := types.MessageMap{
		"name": {
			"required": "Please enter your name",
		},
		"firstname": {"required": "Please enter your name"},
		"lastname":  {"required": "Please enter your last name"},
		"contact": {
			"required": "Please enter your contact",
			"email":    "Not valid email address"},
	}
	return ReturnedAnswerRule, ReturnedAnswerMessages
}
