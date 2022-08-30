package services

import (
	"context"

	"example.com/m/v2/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionServiceImpl struct {
	questionCollection *mongo.Collection
	ctx                context.Context
}

func NewQuestionServices(questionCollection *mongo.Collection, ctx context.Context) *QuestionServiceImpl {
	return &QuestionServiceImpl{
		questionCollection: questionCollection,
		ctx:                ctx,
	}
}

func (u *QuestionServiceImpl) CreateQuestion(Question *models.Question) error {
	_, err := u.questionCollection.InsertOne(u.ctx, Question)
	return err
}

func (u *QuestionServiceImpl) UpdateQuestion(Question *models.Question) error {
	return nil
}

func (u *QuestionServiceImpl) DeleteQuestion(Question *models.Question) error {
	return nil

}

func (u *QuestionServiceImpl) GetQuestion(Question *models.Question) error {
	return nil
}
