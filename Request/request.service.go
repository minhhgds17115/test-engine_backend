package request

import "context"

type RequestServices interface {
	ValidationTest(ctx context.Context) (bool, error)
	CandidateInformationValidate(ctx context.Context) (bool, error)
	ReturnedAnswerValidate(ctx context.Context) (bool, error)
}
