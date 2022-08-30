package services

import (
	"context"

	"example.com/m/v2/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnswerServiceImpl struct {
	answerColllection *mongo.Collection
	ctx               context.Context
}

func NewAnswerServices(answerColllection *mongo.Collection, ctx context.Context) *AnswerServiceImpl {
	return &AnswerServiceImpl{
		answerColllection: answerColllection,
		ctx:               ctx,
	}
}

func (u *AnswerServiceImpl) createAnswer(Question *models.Question) error {
	_, err := u.answerColllection.InsertOne(u.ctx, Question)
	return err
}

func (u *AnswerServiceImpl) deleteAnswer(Question *models.Question) error {
	return nil
}

func (u *AnswerServiceImpl) updateAnswer(Question *models.Question) error {
	return nil

}

func (u *AnswerServiceImpl) getAnswer(Question *models.Question) error {
	return nil
}
